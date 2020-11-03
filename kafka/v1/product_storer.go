package v1

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aidarkhanov/nanoid/v2"
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
func (s *ProductStorer) Store(ctx context.Context, reqID string, pp ...atlant.Product) (err error) {
	if len(pp) == 0 {
		return nil
	}

	if err = s.producer.BeginTransaction(ctx); err != nil {
		return err
	}

	defer func(ctx context.Context, err *error, logger *zap.Logger) {
		if *err == nil {
			return
		}

		if abortErr := s.producer.AbortTransaction(ctx); abortErr != nil {
			logger.Error("error while aborting transaction: ", zap.Error(abortErr))
		}
	}(ctx, &err, s.logger)

	for _, p := range pp {
		var (
			pb = struct {
				CreatedAt int64   `json:"created_at"`
				Price     float64 `json:"price"`
				Name      string  `json:"name"`
				ChangeID  string  `json:"change_id"`
				RequestID string  `json:"request_id"`
			}{
				CreatedAt: p.CreatedAt().UnixNano(),
				Price:     p.Price(),
				Name:      p.Name(),
				RequestID: reqID,
			}

			bb = &bytes.Buffer{}
		)

		if pb.ChangeID, err = nanoid.New(); err != nil {
			return err
		}

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
