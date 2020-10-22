package v1

import (
	"context"
	"io"

	"github.com/stretchr/testify/mock"
)

type Producer interface {
	InitTransactions(ctx context.Context) (err error)
	BeginTransaction(ctx context.Context) (err error)
	AbortTransaction(ctx context.Context) (err error)
	Produce(ctx context.Context, msg io.Reader) (err error)
	CommitTransaction(ctx context.Context) (err error)
}

type MockProducer struct {
	mock.Mock
}

func (p *MockProducer) InitTransactions(ctx context.Context) (err error) {
	return p.Called(ctx).Error(0)
}

func (p *MockProducer) BeginTransaction(ctx context.Context) (err error) {
	return p.Called(ctx).Error(0)
}

func (p *MockProducer) AbortTransaction(ctx context.Context) (err error) {
	return p.Called(ctx).Error(0)
}

func (p *MockProducer) Produce(ctx context.Context, msg io.Reader) (err error) {
	return p.Called(ctx, msg).Error(0)
}

func (p *MockProducer) CommitTransaction(ctx context.Context) (err error) {
	return p.Called(ctx).Error(0)
}
