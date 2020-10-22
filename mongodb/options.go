package mongodb

//
type Option interface {
	//
	apply(c *Client)
}

type mongoClientOptionFunc func(c *Client)

func (f mongoClientOptionFunc) apply(c *Client) {
	f(c)
}

//
func WithURI(uri string) Option {
	return mongoClientOptionFunc(func(c *Client) {
		c.uri = uri
	})
}

//
func WithDatabase(db string) Option {
	return mongoClientOptionFunc(func(c *Client) {
		c.db = db
	})
}

//
func WithCollection(col string) Option {
	return mongoClientOptionFunc(func(c *Client) {
		c.col = col
	})
}
