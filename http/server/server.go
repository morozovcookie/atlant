package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//
type Server struct {
	//
	host string

	//
	router chi.Router

	//
	srv *http.Server
}

//
func New(host string, options ...Option) (s *Server) {
	s = &Server{
		router: chi.NewRouter(),

		host: host,
	}

	// fork gin-zap and implement it for chi
	s.router.Use(middleware.Logger, middleware.RequestID, middleware.Recoverer)

	s.srv = &http.Server{
		Addr:              s.host,
		Handler:           s.router,
		ReadTimeout:       DefaultReadTimeout,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		WriteTimeout:      DefaultWriteTimeout,
		IdleTimeout:       DefaultIdleTimeout,
		MaxHeaderBytes:    DefaultMaxHeaderBytes,
	}

	for _, opt := range options {
		opt.apply(s)
	}

	return s
}

//
func (s *Server) Start() (err error) {
	if err = s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

//
func (s *Server) Stop(ctx context.Context) (err error) {
	return s.srv.Shutdown(ctx)
}
