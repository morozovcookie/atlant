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
type ProductProcessor struct {
	//
	storage atlant.ProductStorage

	//
	logger *zap.Logger
}

//
func NewProductProcessor(ps atlant.ProductStorage, logger *zap.Logger) *ProductProcessor {
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
		CreatedAt int64   `json:"created_at"`
		Price     float64 `json:"price"`
		Name      string  `json:"name"`
		ChangeID  string  `json:"change_id"`
		RequestID string  `json:"request_id"`
	}{}

	// I think that this should be error, but message should committed, because we never can be unmarshal bad message.
	if err = json.NewDecoder(r).Decode(rp); err != nil {
		return kafka.ErrDecodeIncomingMessage
	}

	p := atlant.NewProduct(rp.Name, rp.Price, time.Unix(0, rp.CreatedAt))

	mp, err := pp.storage.GetByProductID(ctx, p.ID())
	if err != nil {
		return err
	}

	if mp == nil {
		return pp.storage.StoreProduct(ctx, p)
	}

	if mp.HasChangeBeenApplied(rp.ChangeID) {
		return nil
	}

	if !mp.HasPriceBeenChanged(p.Price()) {
		return nil
	}

	mp.ApplyChange(&atlant.ProductChanging{
		CreatedAt: time.Unix(0, rp.CreatedAt),
		OldPrice:  mp.Price(),
		NewPrice:  rp.Price,
		RequestID: rp.RequestID,
		ChangeID:  rp.ChangeID,
	})

	return pp.storage.StoreProduct(ctx, mp)
}
