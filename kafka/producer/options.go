package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//
type Option interface {
	apply(p *Producer)
}

type producerOptionFunc func(p *Producer)

func (f producerOptionFunc) apply(p *Producer) {
	f(p)
}

//
func WithServers(servers []string) Option {
	return producerOptionFunc(func(p *Producer) {
		p.servers = servers
	})
}

//
func WithTopic(topic string) Option {
	return producerOptionFunc(func(p *Producer) {
		p.topic = topic
	})
}

//
const DefaultPartition = kafka.PartitionAny

//
func WithPartition(partition int32) Option {
	return producerOptionFunc(func(p *Producer) {
		p.partition = partition
	})
}

const (
	AcknowledgementNoWait     int = 0
	AcknowledgementWaitLeader int = 1
	AcknowledgementWaitAll    int = -1
)

const DefaultAcknowledgement = AcknowledgementWaitLeader

//
func WithAcknowledgement(acks int) Option {
	return producerOptionFunc(func(p *Producer) {
		p.acks = acks
	})
}

//
func WithTransactionalID(id string) Option {
	return producerOptionFunc(func(p *Producer) {
		p.transactionalID = id
	})
}

const (
	IdempotenceDisabledState bool = false
	IdempotenceEnabledState  bool = true
)

const DefaultIdempotenceState = IdempotenceDisabledState

//
func WithIdempotenceState(state bool) Option {
	return producerOptionFunc(func(p *Producer) {
		p.idempotenceState = state
	})
}

const DefaultMaxInFlightRequestsPerConnection = 5

//
func WithMaxInFlightRequestsPerConnection(conns int) Option {
	return producerOptionFunc(func(p *Producer) {
		p.maxInFlightRequestsPerConnection = conns
	})
}

const DefaultRetries = 10000000

//
func WithRetries(retries int) Option {
	return producerOptionFunc(func(p *Producer) {
		p.retries = retries
	})
}

const (
	CompressionTypeNone   string = "none"
	CompressionTypeGzip   string = "gzip"
	CompressionTypeSnappy string = "snappy"
	CompressionTypeLz4    string = "lz4"
	CompressionTypeZstd   string = "zstd"
)

const DefaultCompressionType string = CompressionTypeNone

//
func WithCompressionType(t string) Option {
	return producerOptionFunc(func(p *Producer) {
		p.compressionType = t
	})
}
