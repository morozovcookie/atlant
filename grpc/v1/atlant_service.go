package v1

import (
	"context"
	stderrors "errors"
	"net/url"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/morozovcookie/atlant"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//
type AtlantService struct {
	//
	fetcher atlant.ProductFetcher

	//
	storer atlant.ProductStorer

	//
	lister atlant.ProductLister

	//
	clock Clock

	//
	logger *zap.Logger
}

//
func NewAtlantService(config AtlantServiceConfig, logger *zap.Logger) *AtlantService {
	return &AtlantService{
		fetcher: config.ProductFetcherInstance(),
		storer:  config.ProductStorerInstance(),
		lister:  config.ProductListerInstance(),
		clock:   config.ClockInstance(),

		logger: logger,
	}
}

//
func (s *AtlantService) Fetch(ctx context.Context, r *FetchRequest) (_ *empty.Empty, err error) {
	reqRecvT := s.clock.NowInUTC()

	u, err := url.Parse(r.Url)
	if err != nil {
		s.logger.Error("parse url error", zap.Error(err))

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Optional: think about different protocols and fetchers factory depends on protocols type

	pp, err := s.fetcher.Fetch(ctx, u, reqRecvT)
	if err != nil {
		s.logger.Error("fetch error", zap.Error(err))

		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = s.storer.Store(ctx, pp...); err != nil {
		s.logger.Error("store error", zap.Error(err))

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

var ErrUnknownSortingDirection = errors.New("unknown sorting direction")

//
func (s *AtlantService) List(ctx context.Context, req *ListRequest) (res *ListResponse, err error) {
	var (
		fromProtocolSortingDirectionToDomainSortingDirectionMap = map[string]atlant.SortingDirection{
			ListRequest_SortingOption_SORTING_OPTION_UNSPECIFIED.String(): atlant.SortingDirectionUnspecified,
			ListRequest_SortingOption_SORTING_OPTION_ASC.String():         atlant.SortingDirectionAsc,
			ListRequest_SortingOption_SORTING_OPTION_DESC.String():        atlant.SortingDirectionDesc,
		}

		start    = atlant.NewStartParameter(req.Start)
		limit    = atlant.NewLimitParameter(req.Limit)
		sortOpts = make(atlant.ProductSortingOptions, len(req.Options))
	)

	for i, opt := range req.Options {
		d, ok := fromProtocolSortingDirectionToDomainSortingDirectionMap[opt.Direction.String()]
		if !ok {
			return nil,
				status.Error(codes.InvalidArgument, ErrUnknownSortingDirection.Error()+": "+opt.Direction.String())
		}

		sortOpts[i] = atlant.NewProductSortingOption(atlant.SortingField(opt.Field), d)
	}

	for _, p := range []interface {
		Validate() (err error)
	}{start, limit, sortOpts} {
		err = p.Validate()

		if stderrors.Is(atlant.ErrInvalidLimitParameterMinValue, err) && req.Limit == 0 {
			s.logger.Warn(`"limit" value less than min value - it will be set to 100`)

			limit = atlant.NewLimitParameter(atlant.MaxLimitParameterValue)

			continue
		}

		if stderrors.Is(atlant.ErrInvalidLimitParameterMaxValue, err) {
			s.logger.Warn(`"limit" value more than max value - it will be set to 100`)

			limit = atlant.NewLimitParameter(atlant.MaxLimitParameterValue)

			continue
		}

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	pp, err := s.lister.List(ctx, start, limit, sortOpts)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res = &ListResponse{
		Products: make([]*ListResponse_Product, len(pp)),
	}

	for i, p := range pp {
		res.Products[i] = &ListResponse_Product{
			Name:        p.Name,
			Price:       p.Price,
			CreatedAt:   p.CreatedAt.UnixNano(),
			UpdatedAt:   p.UpdatedAt.UnixNano(),
			UpdateCount: int32(p.UpdateCount),
		}
	}

	return res, nil
}
