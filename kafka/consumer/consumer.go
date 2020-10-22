package consumer

import (
	"bytes"
	"context"
	"errors"
	confluent "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/morozovcookie/atlant/kafka"
	"go.uber.org/zap"
	"io"
	"strings"
)

//
type Consumer struct {
	//
	topic string

	//
	groupID string

	//
	autoOffsetReset string

	//
	isolationLevel string

	//
	clientID string

	//
	servers []string

	//
	kc *confluent.Consumer

	//
	logger *zap.Logger

	//
	ctx context.Context

	//
	cancel context.CancelFunc
}

//
func New(ctx context.Context, logger *zap.Logger, opts ...Option) (c *Consumer, err error) {
	c = &Consumer{
		autoOffsetReset: DefaultAutoOffsetReset,
		isolationLevel:  DefaultIsolationLevel,

		logger: logger,
	}

	c.ctx, c.cancel = context.WithCancel(ctx)

	for _, opt := range opts {
		opt.apply(c)
	}

	c.kc, err = confluent.NewConsumer(&confluent.ConfigMap{
		"enable.auto.commit": false,
		"group.id":           c.groupID,
		"auto.offset.reset":  c.autoOffsetReset,
		"isolation.level":    c.isolationLevel,
		"client.id":          c.clientID,
		"bootstrap.servers":  strings.Join(c.servers, ","),
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

//
func (c *Consumer) Close() (err error) {
	c.cancel()

	return c.kc.Close()
}

//
func (c *Consumer) Subscribe(cb func(r io.Reader) (err error)) (err error) {
	if err = c.kc.Subscribe(c.topic, nil); err != nil {
		return err
	}

	for {
		select {
		case <-c.ctx.Done():
			return nil
		default:
		}

		msg, err := c.kc.ReadMessage(0)
		if err != nil {
			if kerr, ok := err.(confluent.Error); ok {
				if kerr.IsFatal() {
					c.logger.Error("consumer fatal error", zap.Error(err))

					return err
				}

				if kerr.Code() == confluent.ErrTimedOut {
					continue
				}
			}

			c.logger.Error("read message error", zap.Error(err))
			continue
		}

		c.logger.Info("got message", zap.String("message", string(msg.Value)))

		if err = cb(bytes.NewReader(msg.Value)); err != nil {
			if !errors.Is(kafka.ErrDecodeIncomingMessage, err) {
				c.logger.Error("process message error", zap.Error(err))

				continue
			}

			c.logger.Error("decode consumed message error", zap.Error(err), zap.Binary("message", msg.Value))
		}

		if _, err = c.kc.CommitMessage(msg); err != nil {
			c.logger.Error("commit message error", zap.Error(err))

			if kerr, ok := err.(confluent.Error); ok && kerr.IsFatal() {
				return err
			}
		}
	}
}
