package atlant

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"time"

	"github.com/pkg/errors"
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

type StartParameter int64

func NewStartParameter(val int64) (p StartParameter) {
	return StartParameter(val)
}

func (p StartParameter) Int64() (val int64) {
	return (int64)(p)
}

const MinStartParameterValue int64 = 0

var ErrInvalidStartParameterValue = errors.New(`"start" value should be greater or equal zero`)

func (p StartParameter) Validate() (err error) {
	if p.Int64() < MinStartParameterValue {
		return ErrInvalidStartParameterValue
	}

	return nil
}

type LimitParameter int64

func NewLimitParameter(val int64) (p LimitParameter) {
	return LimitParameter(val)
}

func (p LimitParameter) Int64() (val int64) {
	return (int64)(p)
}

const (
	MinLimitParameterValue int64 = 1
	MaxLimitParameterValue int64 = 100
)

var (
	ErrInvalidLimitParameterMinValue = errors.New(`"limit" value should be greater or equal 1`)
	ErrInvalidLimitParameterMaxValue = errors.New(`"limit" value should be less or equal 100`)
)

func (p LimitParameter) Validate() (err error) {
	if p.Int64() < MinLimitParameterValue {
		return ErrInvalidLimitParameterMinValue
	}

	if p.Int64() > MaxLimitParameterValue {
		return ErrInvalidLimitParameterMaxValue
	}

	return nil
}

type SortingField string

func (f SortingField) String() (s string) {
	return string(f)
}

var (
	ErrUnknownField                = errors.New("unknown field")
	ErrFieldNotAvailableForSorting = errors.New("field not available for sorting")
)

func (f SortingField) Validate() (err error) {
	var (
		productFieldsMap = map[string]struct{}{
			"name":         {},
			"price":        {},
			"created_at":   {},
			"updated_at":   {},
			"update_count": {},
		}

		availableSortingFieldsMap = map[string]struct{}{
			"name":         {},
			"price":        {},
			"created_at":   {},
			"updated_at":   {},
			"update_count": {},
		}
	)

	if _, ok := productFieldsMap[f.String()]; !ok {
		return errors.WithMessage(ErrUnknownField, f.String())
	}

	if _, ok := availableSortingFieldsMap[f.String()]; !ok {
		return errors.WithMessage(ErrFieldNotAvailableForSorting, f.String())
	}

	return nil
}

const (
	SortingDirectionUnspecified SortingDirection = "UNSPECIFIED"
	SortingDirectionAsc         SortingDirection = "ASC"
	SortingDirectionDesc        SortingDirection = "DESC"
)

type SortingDirection string

func (d SortingDirection) String() (s string) {
	return string(d)
}

func (d *SortingDirection) Validate() {
	if *d == SortingDirectionUnspecified {
		*d = SortingDirectionAsc
	}
}

//
type ProductSortingOption struct {
	//
	Field SortingField

	//
	Direction SortingDirection
}

func NewProductSortingOption(f SortingField, d SortingDirection) ProductSortingOption {
	return ProductSortingOption{
		Field:     f,
		Direction: d,
	}
}

func (opt ProductSortingOption) Validate() (err error) {
	opt.Direction.Validate()

	return opt.Field.Validate()
}

type ProductSortingOptions []ProductSortingOption

func (opts ProductSortingOptions) Validate() (err error) {
	for _, opt := range opts {
		if err = opt.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// ProductLister retrieve products list.
type ProductLister interface {
	// List retrieve products list with options, like limit or sorting.
	List(ctx context.Context,
		start StartParameter,
		limit LimitParameter,
		opts ProductSortingOptions) (pp []Product, err error)
}

type MockProductLister struct {
	mock.Mock
}

func (pl *MockProductLister) List(
	ctx context.Context,
	start StartParameter,
	limit LimitParameter,
	opts ProductSortingOptions,
) (
	pp []Product,
	err error,
) {
	args := pl.Called(ctx, start, limit, opts)

	return args.Get(0).([]Product), args.Error(1)
}

//
type ProductStorage interface {
	ProductStorer

	//
	GetByProductID(ctx context.Context, productID string) (p *Product, err error)
}

type MockProductStorage struct {
	mock.Mock
}

func (ps *MockProductStorage) Store(ctx context.Context, pp ...Product) (err error) {
	return ps.Called(ctx, pp).Error(0)
}

func (ps *MockProductStorage) GetByProductID(ctx context.Context, productID string) (p *Product, err error) {
	args := ps.Called(ctx, productID)

	return args.Get(0).(*Product), args.Error(1)
}
