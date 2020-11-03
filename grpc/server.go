package grpc

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//
type Server struct {
	//
	address string

	//
	srv *grpc.Server

	//
	logger *zap.Logger
}

//
func NewServer(address string, logger *zap.Logger, opts ...Option) (s *Server) {
	s = &Server{
		address: address,

		srv: grpc.NewServer(),

		logger: logger,
	}

	for _, opt := range opts {
		opt.apply(s)
	}

	return s
}

//
func (s *Server) ListenAndServe() (err error) {
	s.logger.Info("starting")

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	return s.srv.Serve(lis)
}

//
func (s *Server) Close() {
	s.srv.GracefulStop()
}
