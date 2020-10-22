package v1

import (
	"context"
	"net/url"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/morozovcookie/atlant"
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
	clock Clock

	//
	logger *zap.Logger
}

//
func NewAtlantService(config AtlantServiceConfig, logger *zap.Logger) *AtlantService {
	return &AtlantService{
		fetcher: config.ProductFetcherInstance(),
		storer:  config.ProductStorerInstance(),
		clock:   config.ClockInstance(),

		logger: logger,
	}
}

//
func (s *AtlantService) Fetch(ctx context.Context, r *FetchRequest) (_ *empty.Empty, err error) {
	reqRecvT := s.clock.NowInUTC()

	u, err := url.Parse(r.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Optional: think about different protocols and fetchers factory depends on protocols type

	pp, err := s.fetcher.Fetch(ctx, u, reqRecvT)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = s.storer.Store(ctx, pp...); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

//
func (s *AtlantService) List(ctx context.Context, req *ListRequest) (res *ListResponse, err error) {
	return nil, nil
}
