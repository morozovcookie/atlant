package v1

import (
	"github.com/morozovcookie/atlant"
)

//
type Container struct {
	//
	ProductFetcher atlant.ProductFetcher

	//
	ProductStorer atlant.ProductStorer

	//
	Clock Clock
}

//
func (c Container) ProductFetcherInstance() (fetcher atlant.ProductFetcher) {
	return c.ProductFetcher
}

//
func (c Container) ProductStorerInstance() (storer atlant.ProductStorer) {
	return c.ProductStorer
}

//
func (c Container) ClockInstance() (clock Clock) {
	return c.Clock
}
