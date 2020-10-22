package v1

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/morozovcookie/atlant"
	"go.uber.org/zap"
)

//
type ProductStorer struct {
	//
	producer Producer

	//
	logger *zap.Logger
}

//
func NewProductStorer(producer Producer, logger *zap.Logger) (ps *ProductStorer) {
	return &ProductStorer{
		producer: producer,
		logger:   logger,
	}
}

//
func (s *ProductStorer) Store(ctx context.Context, pp ...atlant.Product) (err error) {
	if len(pp) == 0 {
		return nil
	}

	if err = s.producer.BeginTransaction(ctx); err != nil {
		return err
	}

	func(ctx context.Context, err error, logger *zap.Logger) {
		if err == nil {
			return
		}

		if abortErr := s.producer.AbortTransaction(ctx); abortErr != nil {
			logger.Error("error while aborting transaction: ", zap.Error(abortErr))
		}
	}(ctx, err, s.logger)

	for _, p := range pp {
		var (
			pb = struct {
				Name      string  `json:"name"`
				Price     float64 `json:"price"`
				CreatedAt int64   `json:"created_at"`
			}{
				Name:      p.Name,
				Price:     p.Price,
				CreatedAt: p.CreatedAt.UnixNano(),
			}

			bb = &bytes.Buffer{}
		)

		if err = json.NewEncoder(bb).Encode(pb); err != nil {
			return err
		}

		if err = s.producer.Produce(ctx, bb); err != nil {
			return err
		}
	}

	if err = s.producer.CommitTransaction(ctx); err != nil {
		return err
	}

	return nil
}
