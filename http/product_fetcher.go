package http

import (
	"context"
	"encoding/csv"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/morozovcookie/atlant"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var ErrTooManyRequests = errors.New("too many requests")

type ProductFetcher struct {
	c Client

	requestTimeout time.Duration
	maxRequests    int

	logger *zap.Logger
}

func NewProductFetcher(c Client, logger *zap.Logger) (f *ProductFetcher) {
	return &ProductFetcher{
		c: c,

		requestTimeout: time.Millisecond * 10,
		maxRequests:    10,

		logger: logger,
	}
}

func (f *ProductFetcher) Fetch(ctx context.Context, u *url.URL, timeMark time.Time) (pp []atlant.Product, err error) {
	var (
		i    int
		resp *http.Response
	)

	for {
		f.logger.Debug("fetch file", zap.Int("attempt", i+1))

		resp, err = f.c.Get(ctx, u.String())
		if err != nil {
			return nil, err
		}

		f.logger.Debug("got response",
			zap.Int("status_code", resp.StatusCode),
			zap.String("status", resp.Status),
			zap.Int64("content_length", resp.ContentLength))

		if fileWasFound(resp) {
			break
		}

		if fileDoesNotExist(resp) {
			return nil, atlant.ErrFileDoesNotExist
		}

		if !shouldTryAgain(resp, f.maxRequests, i) {
			return nil, ErrTooManyRequests
		}

		f.logger.Debug("wait before another call", zap.Duration("timeout", f.requestTimeout))

		i++
		<-time.After(f.requestTimeout)
	}

	defer func(closer io.Closer, logger *zap.Logger) {
		if closeErr := closer.Close(); closeErr != nil {
			logger.Error("close response error", zap.Error(closeErr))
			err = closeErr
		}
	}(resp.Body, f.logger)

	r := csv.NewReader(resp.Body)
	r.Comma = ';'

	ss, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(ss) == 0 {
		return nil, nil
	}

	pp = make([]atlant.Product, 0, len(ss))

	for i, s := range ss {
		f.logger.Debug("parse line",
			zap.Int("line number", i),
			zap.Strings("line", s))

		price, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			return nil, err
		}

		pp = append(pp, *(atlant.NewProduct(s[0], price, timeMark)))
	}

	return pp, nil
}

func fileDoesNotExist(r *http.Response) bool {
	return r.StatusCode == http.StatusNotFound
}

func fileWasFound(r *http.Response) bool {
	return r.StatusCode == http.StatusOK
}

func shouldTryAgain(r *http.Response, max, i int) bool {
	return r.StatusCode >= http.StatusInternalServerError &&
		r.StatusCode <= http.StatusNetworkAuthenticationRequired &&
		i < max
}
