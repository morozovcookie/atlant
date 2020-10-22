package v1

import (
	"github.com/morozovcookie/atlant"
	"github.com/stretchr/testify/mock"
)

//
type AtlantServiceConfig interface {
	//
	ProductFetcherInstance() atlant.ProductFetcher

	//
	ProductStorerInstance() atlant.ProductStorer

	//
	ProductListerInstance() atlant.ProductLister

	//
	ClockInstance() Clock
}

type MockAtlantServiceConfig struct {
	mock.Mock
}

func (c *MockAtlantServiceConfig) ProductFetcherInstance() atlant.ProductFetcher {
	return c.Called().Get(0).(atlant.ProductFetcher)
}

func (c *MockAtlantServiceConfig) ProductStorerInstance() atlant.ProductStorer {
	return c.Called().Get(0).(atlant.ProductStorer)
}

func (c *MockAtlantServiceConfig) ProductListerInstance() atlant.ProductLister {
	return c.Called().Get(0).(atlant.ProductLister)
}

func (c *MockAtlantServiceConfig) ClockInstance() Clock {
	return c.Called().Get(0).(Clock)
}
