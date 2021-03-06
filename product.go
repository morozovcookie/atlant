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

// ProductChanging represent record about products price changing.
type ProductChanging struct {
	// OldPrice is a product price before changing
	OldPrice float64 `bson:"old_price"`

	// NewPrice is a product price after changing
	NewPrice float64 `bson:"new_price"`

	// RequestID is a client request identifier
	RequestID string `bson:"request_id"`

	// ChangeID is a changing identifier
	ChangeID string `bson:"change_id"`

	// CreatedAt is a time of product price changing
	CreatedAt time.Time `bson:"created_at"`
}

// Product represent product record from external source.
type Product struct {
	// updateCount represent how many times product was modified
	updateCount int

	// price is a current products price
	price float64

	// name is products name
	name string

	// createdAt is a products created time
	createdAt time.Time

	// updatedAt is a products price time updated
	updatedAt time.Time

	// changeHistoryDataIndex
	changeHistoryDataIndex map[string]*ProductChanging

	// changeHistoryData is a set of all changes of current product
	// maybe the best option will be replace slice with list? and replace index with some structure too.
	changeHistoryData []ProductChanging
}

//
func NewProduct(name string, price float64, createdAt time.Time) (p *Product) {
	return &Product{
		updateCount:            0,
		price:                  price,
		name:                   name,
		createdAt:              createdAt,
		updatedAt:              time.Unix(0, 0),
		changeHistoryDataIndex: make(map[string]*ProductChanging),
	}
}

//
func (p Product) UpdateCount() (c int) {
	return p.updateCount
}

//
func (p Product) Price() (price float64) {
	return p.price
}

//
func (p Product) Name() (name string) {
	return p.name
}

//
func (p Product) CreatedAt() (t time.Time) {
	return p.createdAt
}

//
func (p Product) UpdatedAt() (t time.Time) {
	return p.updatedAt
}

//
func (p Product) ChangeHistory() (ch []ProductChanging) {
	return p.changeHistoryData
}

//
func (p Product) HasChangeBeenApplied(changeID string) (hasBeenApplied bool) {
	_, hasBeenApplied = p.changeHistoryDataIndex[changeID]

	return hasBeenApplied
}

//
func (p Product) HasPriceBeenChanged(price float64) (hasPriceChanged bool) {
	return p.price != price
}

//
func (p *Product) ApplyChange(c *ProductChanging) {
	p.updateCount++
	p.price = c.NewPrice
	p.updatedAt = c.CreatedAt

	p.changeHistoryData = append(p.changeHistoryData, ProductChanging{
		OldPrice:  c.OldPrice,
		NewPrice:  c.NewPrice,
		RequestID: c.RequestID,
		ChangeID:  c.ChangeID,
		CreatedAt: c.CreatedAt,
	})
	p.changeHistoryDataIndex[c.ChangeID] = &p.changeHistoryData[len(p.changeHistoryData)-1]
}

//
func (p Product) ID() (id string) {
	hs := sha256.Sum256(append([]byte{}, p.name...))

	return hex.EncodeToString(hs[:])
}

var ErrFileDoesNotExist = errors.New("file with specified URL does not exist")

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
	Store(ctx context.Context, reqID string, pp ...Product) (err error)
}

type MockProductStorer struct {
	mock.Mock
}

func (ps *MockProductStorer) Store(ctx context.Context, reqID string, pp ...Product) (err error) {
	return ps.Called(ctx, pp).Error(0)
}

//
type StartParameter int64

//
func NewStartParameter(val int64) (p StartParameter) {
	return StartParameter(val)
}

//
func (p StartParameter) Int64() (val int64) {
	return (int64)(p)
}

//
const MinStartParameterValue int64 = 0

// ErrInvalidStartParameterValue raised when start value less than zero.
var ErrInvalidStartParameterValue = errors.New(`"start" value should be greater or equal zero`)

//
func (p StartParameter) Validate() (err error) {
	if p.Int64() < MinStartParameterValue {
		return ErrInvalidStartParameterValue
	}

	return nil
}

//
type LimitParameter int64

//
func NewLimitParameter(val int64) (p LimitParameter) {
	return LimitParameter(val)
}

//
func (p LimitParameter) Int64() (val int64) {
	return (int64)(p)
}

const (
	//
	MinLimitParameterValue int64 = 1

	//
	MaxLimitParameterValue int64 = 100
)

var (
	// ErrInvalidLimitParameterMinValue raise when limit value less than 1.
	ErrInvalidLimitParameterMinValue = errors.New(`"limit" value should be greater or equal 1`)

	// ErrInvalidLimitParameterMaxValue raise when limit value greater than maximum value (100).
	ErrInvalidLimitParameterMaxValue = errors.New(`"limit" value should be less or equal 100`)
)

//
func (p LimitParameter) Validate() (err error) {
	if p.Int64() < MinLimitParameterValue {
		return ErrInvalidLimitParameterMinValue
	}

	if p.Int64() > MaxLimitParameterValue {
		return ErrInvalidLimitParameterMaxValue
	}

	return nil
}

//
type SortingField string

//
func (f SortingField) String() (s string) {
	return string(f)
}

var (
	// ErrUnknownField raise when unknown field was passed in sorting parameters.
	ErrUnknownField = errors.New("unknown field")

	// ErrFieldNotAvailableForSorting when field which was passed through sorting parameters was recognized but it not
	// available for sorting.
	ErrFieldNotAvailableForSorting = errors.New("field not available for sorting")
)

//
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
	//
	SortingDirectionUnspecified SortingDirection = "UNSPECIFIED"

	//
	SortingDirectionAsc SortingDirection = "ASC"

	//
	SortingDirectionDesc SortingDirection = "DESC"
)

//
type SortingDirection string

//
func (d SortingDirection) String() (s string) {
	return string(d)
}

//
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

//
func NewProductSortingOption(f SortingField, d SortingDirection) ProductSortingOption {
	return ProductSortingOption{
		Field:     f,
		Direction: d,
	}
}

//
func (opt ProductSortingOption) Validate() (err error) {
	opt.Direction.Validate()

	return opt.Field.Validate()
}

//
type ProductSortingOptions []ProductSortingOption

//
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
	//
	StoreProduct(ctx context.Context, p *Product) (err error)

	//
	GetByProductID(ctx context.Context, productID string) (p *Product, err error)
}

type MockProductStorage struct {
	mock.Mock
}

func (ps *MockProductStorage) StoreProduct(ctx context.Context, p *Product) (err error) {
	return ps.Called(ctx, p).Error(0)
}

func (ps *MockProductStorage) GetByProductID(ctx context.Context, productID string) (p *Product, err error) {
	args := ps.Called(ctx, productID)

	return args.Get(0).(*Product), args.Error(1)
}
