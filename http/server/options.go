package server

import (
	"net/http"
	"time"
)

// Option represent any option for configuring HTTP server.
type Option interface {
	apply(s *Server)
}

type serverOptionFunc func(s *Server)

func (f serverOptionFunc) apply(s *Server) {
	f(s)
}

// DefaultReadTimeout is the default duration for reading the entire
// request, including the body.
const DefaultReadTimeout = time.Millisecond * 100

// WithReadTimeout set the maximum duration for reading the entire
// request, including the body.
func WithReadTimeout(timeout time.Duration) Option {
	return serverOptionFunc(func(s *Server) {
		s.srv.ReadTimeout = timeout
	})
}

// DefaultReadHeaderTimeout is the default amount of time allowed to read
// request headers.
const DefaultReadHeaderTimeout = time.Millisecond * 100

// WithReadHeaderTimeout set the amount of time allowed to read
// request headers.
func WithReadHeaderTimeout(timeout time.Duration) Option {
	return serverOptionFunc(func(s *Server) {
		s.srv.ReadHeaderTimeout = timeout
	})
}

// DefaultWriteTimeout is the default duration before timing out
// writes of the response.
const DefaultWriteTimeout = time.Millisecond * 100

// WithWriteTimeout set the maximum duration before timing out
// writes of the response.
func WithWriteTimeout(timeout time.Duration) Option {
	return serverOptionFunc(func(s *Server) {
		s.srv.WriteTimeout = timeout
	})
}

// DefaultIdleTimeout is the default amount of time to wait for the
// next request when keep-alives are enabled.
const DefaultIdleTimeout = time.Millisecond * 100

// WithIdleTimeout set the maximum amount of time to wait for the
// next request when keep-alives are enabled.
func WithIdleTimeout(timeout time.Duration) Option {
	return serverOptionFunc(func(s *Server) {
		s.srv.IdleTimeout = timeout
	})
}

// DefaultMaxHeaderBytes controls the default number of bytes the
// server will read parsing the request header's keys and
// values, including the request line.
const DefaultMaxHeaderBytes = http.DefaultMaxHeaderBytes

// WithMaxHeaderBytes set the maximum number of bytes the
// server will read parsing the request header's keys and
// values, including the request line.
func WithMaxHeaderBytes(bytes int) Option {
	return serverOptionFunc(func(s *Server) {
		s.srv.MaxHeaderBytes = bytes
	})
}

// WithHandler append specified sub-router to server's super-router
// by pattern. It is allow to hide server's router interface from
// public usage.
func WithHandler(pattern string, handler http.Handler) Option {
	return serverOptionFunc(func(s *Server) {
		s.router.Mount(pattern, handler)
	})
}
