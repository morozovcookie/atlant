package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	uri string
	db  string
	col string

	client *mongo.Client
}

const (
	connectTimeout = 5 * time.Second
)

func NewClient(opts ...Option) (c *Client) {
	c = &Client{}

	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (c *Client) Connect(ctx context.Context) (err error) {
	connCtx, connCancel := context.WithTimeout(ctx, connectTimeout)
	defer connCancel()

	if c.client, err = mongo.NewClient(options.Client().ApplyURI(c.uri)); err != nil {
		return err
	}

	if err = c.client.Connect(connCtx); err != nil {
		return err
	}

	return c.client.Ping(connCtx, nil)
}

func (c *Client) Close(ctx context.Context) (err error) {
	return c.client.Disconnect(ctx)
}

func (c *Client) Collection() (mc *mongo.Collection) {
	return c.client.Database(c.db).Collection(c.col)
}
