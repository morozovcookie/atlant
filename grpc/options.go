package grpc

import (
	"google.golang.org/grpc"
)

//
type Option interface {
	apply(s *Server)
}

type serverOptionFunc func(s *Server)

func (f serverOptionFunc) apply(s *Server) {
	f(s)
}

//
func WithServiceRegistrator(reg func(gs *grpc.Server)) Option {
	return serverOptionFunc(func(s *Server) {
		reg(s.gs)
	})
}
