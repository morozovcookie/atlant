package http

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

//
type Client interface {
	//
	Get(ctx context.Context, url string) (resp *http.Response, err error)
}

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	args := c.Called(ctx, url)

	return args.Get(0).(*http.Response), args.Error(1)
}
