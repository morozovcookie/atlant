package v1

import (
	"github.com/morozovcookie/atlant"
)

// TODO: this is weird - I should move this file to another place

//
type Container struct {
	//
	ProductFetcher atlant.ProductFetcher

	//
	ProductStorer atlant.ProductStorer

	//
	ProductLister atlant.ProductLister

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
func (c Container) ProductListerInstance() (lister atlant.ProductLister) {
	return c.ProductLister
}

//
func (c Container) ClockInstance() (clock Clock) {
	return c.Clock
}
