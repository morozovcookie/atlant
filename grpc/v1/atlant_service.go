package v1

import (
	"context"
	stderrors "errors"
	"fmt"
	"net/url"

	"github.com/aidarkhanov/nanoid/v2"
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
	reqID, err := nanoid.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.logger.Info("receive Fetch request", zap.String("request", fmt.Sprintf("%+v", r)))

	u, err := url.Parse(r.Url)
	if err != nil {
		s.logger.Error("parse url error", zap.Error(err))

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Optional: think about different protocols and fetchers factory depends on protocols type

	pp, err := s.fetcher.Fetch(ctx, u, s.clock.NowInUTC())
	if err != nil {
		if stderrors.Is(atlant.ErrFileDoesNotExist, err) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		s.logger.Error("fetch error", zap.Error(err))

		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = s.storer.Store(ctx, reqID, pp...); err != nil {
		s.logger.Error("store error", zap.Error(err))

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

var ErrUnknownSortingDirection = errors.New("unknown sorting direction")

//
func (s *AtlantService) List(ctx context.Context, req *ListRequest) (res *ListResponse, err error) {
	s.logger.Info("receive List request", zap.String("request", fmt.Sprintf("%+v", req)))

	var (
		start = atlant.NewStartParameter(req.Start)
		limit = atlant.NewLimitParameter(req.Limit)
	)

	sortOpts, err := initSortingOptions(req.Options)
	if err != nil {
		s.logger.Error("init sorting options error", zap.Error(err))

		return nil, err
	}

	if err = validateListRequestParameters(start, &limit, req.Limit, sortOpts, s.logger); err != nil {
		s.logger.Error("validate list request parameters", zap.Error(err))

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pp, err := s.lister.List(ctx, start, limit, sortOpts)
	if err != nil {
		s.logger.Error("list request error", zap.Error(err))

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res = &ListResponse{
		Products: make([]*ListResponse_Product, len(pp)),
	}

	for i, p := range pp {
		res.Products[i] = &ListResponse_Product{
			Name:        p.Name(),
			Price:       p.Price(),
			CreatedAt:   p.CreatedAt().UnixNano(),
			UpdatedAt:   p.UpdatedAt().UnixNano(),
			UpdateCount: int32(p.UpdateCount()),
		}
	}

	return res, nil
}

//
func initSortingOptions(reqOpts []*ListRequest_SortingOption) (opts atlant.ProductSortingOptions, err error) {
	opts = make(atlant.ProductSortingOptions, len(reqOpts))

	fromProtocolSortingDirectionToDomainSortingDirectionMap := map[string]atlant.SortingDirection{
		ListRequest_SortingOption_SORTING_OPTION_UNSPECIFIED.String(): atlant.SortingDirectionUnspecified,
		ListRequest_SortingOption_SORTING_OPTION_ASC.String():         atlant.SortingDirectionAsc,
		ListRequest_SortingOption_SORTING_OPTION_DESC.String():        atlant.SortingDirectionDesc,
	}

	for i, opt := range reqOpts {
		d, ok := fromProtocolSortingDirectionToDomainSortingDirectionMap[opt.Direction.String()]
		if !ok {
			return nil,
				status.Error(codes.InvalidArgument, ErrUnknownSortingDirection.Error()+": "+opt.Direction.String())
		}

		opts[i] = atlant.NewProductSortingOption(atlant.SortingField(opt.Field), d)
	}

	return opts, nil
}

//
func validateListRequestParameters(
	start atlant.StartParameter,
	limit *atlant.LimitParameter,
	requestedLimit int64,
	opts atlant.ProductSortingOptions,
	logger *zap.Logger,
) (
	err error,
) {
	for _, p := range []interface {
		Validate() (err error)
	}{start, *limit, opts} {
		err = p.Validate()

		if stderrors.Is(atlant.ErrInvalidLimitParameterMinValue, err) && requestedLimit == 0 {
			logger.Warn(`"limit" value less than min value - it will be set to 100`)

			*limit = atlant.NewLimitParameter(atlant.MaxLimitParameterValue)

			continue
		}

		if stderrors.Is(atlant.ErrInvalidLimitParameterMaxValue, err) {
			logger.Warn(`"limit" value more than max value - it will be set to 100`)

			*limit = atlant.NewLimitParameter(atlant.MaxLimitParameterValue)

			continue
		}

		if err != nil {
			logger.Error("validate request parameters error", zap.Error(err))

			return err
		}
	}

	return nil
}
