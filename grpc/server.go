package grpc

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServerCredentials struct {
	CrtPath string
	KeyPath string
}

//
type Server struct {
	//
	gs *grpc.Server

	//
	host string

	//
	logger *zap.Logger

	//
	creds *ServerCredentials
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
	s.gs = grpc.NewServer()

	if s.creds != nil {
		creds, err := credentials.NewServerTLSFromFile(s.creds.CrtPath, s.creds.KeyPath)
		if err != nil {
			return err
		}

		s.gs = grpc.NewServer(grpc.Creds(creds))
	}

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
