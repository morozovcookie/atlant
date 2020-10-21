package client

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

//
type Client struct {
	//
	t *http.Transport

	//
	c *http.Client

	//
	logger *zap.Logger
}

//
func New(logger *zap.Logger, opts ...Option) (c *Client) {
	c = &Client{
		t: &http.Transport{
			MaxIdleConns:           DefaultMaxIdleConns,
			MaxIdleConnsPerHost:    DefaultMaxIdleConnsPerHost,
			MaxConnsPerHost:        DefaultMaxConnsPerHost,
			IdleConnTimeout:        DefaultIdleConnTimeout,
			ResponseHeaderTimeout:  DefaultResponseHeaderTimeout,
			ExpectContinueTimeout:  DefaultExpectContinueTimeout,
			MaxResponseHeaderBytes: DefaultMaxResponseHeaderBytes,
			WriteBufferSize:        DefaultWriteBufferSize,
			ReadBufferSize:         DefaultReadBufferSize,
		},

		c: &http.Client{},

		logger: logger,
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	c.c.Transport = c.t

	return c
}

//
func (c *Client) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.c.Do(req)
}
