package v1

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/morozovcookie/atlant"
	"github.com/morozovcookie/atlant/kafka"
	"go.uber.org/zap"
)

//
type ProductStorage interface {
	//
	GetByID(ctx context.Context, id string) (p *atlant.Product, err error)

	//
	Store(ctx context.Context, pp ...atlant.Product) (err error)
}

//
type ProductProcessor struct {
	//
	storage ProductStorage

	//
	logger *zap.Logger
}

//
func NewProductProcessor(ps ProductStorage, logger *zap.Logger) *ProductProcessor {
	return &ProductProcessor{
		storage: ps,

		logger: logger,
	}
}

//
func (pp *ProductProcessor) ProcessProduct(r io.Reader) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rp := &struct {
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		CreatedAt int64   `json:"created_at"`
		UpdatedAt int64   `json:"updated_at"`
	}{}

	// I think that this should be error, but message should committed, because we never can be unmarshal bad message.
	if err = json.NewDecoder(r).Decode(rp); err != nil {
		return kafka.ErrDecodeIncomingMessage
	}

	p := &atlant.Product{
		Name:      rp.Name,
		Price:     rp.Price,
		CreatedAt: time.Unix(0, rp.CreatedAt),
		UpdatedAt: time.Unix(0, rp.UpdatedAt),
	}

	mp, err := pp.storage.GetByID(ctx, p.ID())
	if err != nil {
		return err
	}

	// if exists and price did not changed -> skip
	if mp != nil && mp.Price == p.Price {
		return nil
	}

	// Update if exist
	if mp != nil {
		mp.Price = p.Price
		mp.UpdatedAt = p.CreatedAt
		mp.UpdateCount++
	}

	if mp == nil {
		mp = p
	}

	return pp.storage.Store(ctx, *mp)
}
