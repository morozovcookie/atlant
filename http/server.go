package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Addr string

	router chi.Router

	srv *http.Server
}

func NewServer(addr string, options ...ServerOption) (server *Server) {
	server = &Server{
		router: chi.NewRouter(),

		Addr: addr,
	}

	// fork gin-zap and implement it for chi
	server.router.Use(middleware.Logger, middleware.RequestID, middleware.Recoverer)

	server.srv = &http.Server{
		Addr:              server.Addr,
		Handler:           server.router,
		ReadTimeout:       DefaultReadTimeout,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		WriteTimeout:      DefaultWriteTimeout,
		IdleTimeout:       DefaultIdleTimeout,
		MaxHeaderBytes:    DefaultMaxHeaderBytes,
	}

	for _, opt := range options {
		opt.apply(server)
	}

	return server
}

func (server *Server) Start() (err error) {
	if err = server.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (server *Server) Stop(ctx context.Context) (err error) {
	return server.srv.Shutdown(ctx)
}
