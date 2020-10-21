package client

import (
	"net/http"
	"time"
)

//
type Option interface {
	apply(c *Client)
}

type clientOptionFunc func(c *Client)

func (f clientOptionFunc) apply(c *Client) {
	f(c)
}

//
const DefaultMaxIdleConns = 100

//
func WithMaxIdleConns(conns int) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.MaxIdleConns = conns
	})
}

//
const DefaultMaxIdleConnsPerHost = http.DefaultMaxIdleConnsPerHost

//
func WithMaxIdleConnsPerHost(conns int) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.MaxIdleConnsPerHost = conns
	})
}

//
const DefaultMaxConnsPerHost = 2

//
func WithMaxConnsPerHost(conns int) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.MaxConnsPerHost = conns
	})
}

//
const DefaultIdleConnTimeout = 90 * time.Second

//
func WithIdleConnTimeout(d time.Duration) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.IdleConnTimeout = d
	})
}

//
const DefaultResponseHeaderTimeout = 1 * time.Second

//
func WithResponseHeaderTimeout(d time.Duration) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.ResponseHeaderTimeout = d
	})
}

//
const DefaultExpectContinueTimeout = 1 * time.Second

//
func WithExpectContinueTimeout(d time.Duration) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.ExpectContinueTimeout = d
	})
}

//
const DefaultMaxResponseHeaderBytes = 4096

//
func WithMaxResponseHeaderBytes(size int64) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.MaxResponseHeaderBytes = size
	})
}

//
const DefaultWriteBufferSize = 4096

//
func WithWriteBufferSize(size int) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.WriteBufferSize = size
	})
}

//
const DefaultReadBufferSize = 4096

//
func WithReadBufferSize(size int) Option {
	return clientOptionFunc(func(c *Client) {
		c.t.ReadBufferSize = size
	})
}
