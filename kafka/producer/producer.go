package producer

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

//
type Producer struct {
	//
	idempotenceState bool

	//
	partition int32

	//
	acks int

	//
	maxInFlightRequestsPerConnection int

	//
	retries int

	//
	topic string

	//
	transactionalID string

	//
	compressionType string

	//
	servers []string

	//
	kp *kafka.Producer

	//
	logger *zap.Logger
}

//
func New(logger *zap.Logger, opts ...Option) (p *Producer, err error) {
	p = &Producer{
		idempotenceState:                 DefaultIdempotenceState,
		partition:                        DefaultPartition,
		acks:                             DefaultAcknowledgement,
		maxInFlightRequestsPerConnection: DefaultMaxInFlightRequestsPerConnection,
		retries:                          DefaultRetries,

		logger: logger,
	}

	for _, opt := range opts {
		opt.apply(p)
	}

	p.kp, err = kafka.NewProducer(&kafka.ConfigMap{
		"enable.idempotence":                    p.idempotenceState,
		"acks":                                  p.acks,
		"max.in.flight.requests.per.connection": p.maxInFlightRequestsPerConnection,
		"retries":                               p.retries,
		"transactional.id":                      p.transactionalID,
		"compression.type":                      p.compressionType,
		"bootstrap.servers":                     strings.Join(p.servers, ","),
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}

//
func (p *Producer) InitTransactions(ctx context.Context) (err error) {
	return p.kp.InitTransactions(ctx)
}

//
func (p *Producer) BeginTransaction(_ context.Context) (err error) {
	return p.kp.BeginTransaction()
}

//
func (p *Producer) AbortTransaction(ctx context.Context) (err error) {
	return p.kp.AbortTransaction(ctx)
}

//
func (p *Producer) Produce(_ context.Context, msg io.Reader) (err error) {
	buf := &bytes.Buffer{}
	if _, err = io.Copy(buf, msg); err != nil {
		return err
	}

	return p.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: p.partition,
		},
		Value: buf.Bytes(),
	}, nil)
}

//
func (p *Producer) CommitTransaction(ctx context.Context) (err error) {
	return p.kp.CommitTransaction(ctx)
}

//
func (p *Producer) Close(_ context.Context) {
	p.kp.Close()
}
