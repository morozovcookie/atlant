package grpc

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//
type Server struct {
	//
	gs *grpc.Server

	//
	host string

	//
	logger *zap.Logger
}

//
func NewServer(host string, logger *zap.Logger, opts ...Option) (s *Server) {
	s = &Server{
		gs:     grpc.NewServer(),
		host:   host,
		logger: logger,
	}

	for _, opt := range opts {
		opt.apply(s)
	}

	return s
}

//
func (s *Server) Start() (err error) {
	lis, err := net.Listen("tcp", s.host)
	if err != nil {
		return err
	}

	return s.gs.Serve(lis)
}

//
func (s *Server) Stop() {
	s.gs.GracefulStop()
}
