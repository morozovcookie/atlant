package http

import (
	"context"
	"encoding/csv"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/morozovcookie/atlant"
	"go.uber.org/zap"
)

type ProductFetcher struct {
	c Client

	logger *zap.Logger
}

func NewProductFetcher(c Client, logger *zap.Logger) *ProductFetcher {
	return &ProductFetcher{
		c: c,

		logger: logger,
	}
}

func (f *ProductFetcher) Fetch(ctx context.Context, u *url.URL, timeMark time.Time) (pp []atlant.Product, err error) {
	resp, err := f.c.Get(ctx, u.String())
	if err != nil {
		return nil, err
	}

	defer func(closer io.Closer, logger *zap.Logger) {
		if closeErr := closer.Close(); closeErr != nil {
			logger.Error("close response error", zap.Error(closeErr))
			err = closeErr
		}
	}(resp.Body, f.logger)

	f.logger.Debug("got response",
		zap.Int("status_code", resp.StatusCode),
		zap.String("status", resp.Status),
		zap.Int64("content_length", resp.ContentLength))

	// TODO: add circuit breaker

	r := csv.NewReader(resp.Body)
	r.Comma = ';'

	ss, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(ss) == 0 {
		return nil, nil
	}

	pp = make([]atlant.Product, len(ss))

	for i, s := range ss {
		f.logger.Debug("parse line",
			zap.Int("line number", i),
			zap.Strings("line", s))

		pp[i] = atlant.Product{
			Name:      s[0],
			CreatedAt: timeMark,
		}

		if pp[i].Price, err = strconv.ParseFloat(s[1], 64); err != nil {
			return nil, err
		}
	}

	return pp, nil
}
