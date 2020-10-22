package atlant

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"time"

	"github.com/stretchr/testify/mock"
)

// Product represent product record from external source.
type Product struct {
	// UpdateCount represent how many times product was modified
	UpdateCount int

	// Price is a current products price
	Price float64

	// Name is products name
	Name string

	// CreatedAt is a products created time
	CreatedAt time.Time

	// UpdatedAt is a products price time updated
	UpdatedAt time.Time
}

func (p Product) ID() (id string) {
	hs := sha256.Sum256(append([]byte{}, p.Name...))

	return hex.EncodeToString(hs[:])
}

// ProductFetcher fetch products list from external resource.
type ProductFetcher interface {
	// Fetch get products list from external source by url.
	Fetch(ctx context.Context, u *url.URL, timeMark time.Time) (pp []Product, err error)
}

type MockProductFetcher struct {
	mock.Mock
}

func (pf *MockProductFetcher) Fetch(ctx context.Context, u *url.URL, timeMark time.Time) (pp []Product, err error) {
	r := pf.Called(ctx, u, timeMark)

	return r.Get(0).([]Product), r.Error(1)
}

// ProductStorer save products list.
type ProductStorer interface {
	// Store saves products collection in storage.
	Store(ctx context.Context, pp ...Product) (err error)
}

type MockProductStorer struct {
	mock.Mock
}

func (ps *MockProductStorer) Store(ctx context.Context, pp ...Product) (err error) {
	return ps.Called(ctx, pp).Error(0)
}

//
type ProductStorage interface {
	ProductStorer

	//
	GetByID(ctx context.Context, id string) (p Product, err error)
}

type MockProductStorage struct {
	mock.Mock
}

func (ps *MockProductStorage) Store(ctx context.Context, pp ...Product) (err error) {
	return ps.Called(ctx, pp).Error(0)
}

func (ps *MockProductStorage) GetByID(ctx context.Context, id string) (p Product, err error) {
	args := ps.Called(ctx, id)

	return args.Get(0).(Product), args.Error(1)
}
