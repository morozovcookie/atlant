package consumer

//
type Option interface {
	apply(c *Consumer)
}

type consumerOptionFunc func(c *Consumer)

func (f consumerOptionFunc) apply(c *Consumer) {
	f(c)
}

//
func WithServers(servers []string) Option {
	return consumerOptionFunc(func(c *Consumer) {
		c.servers = servers
	})
}

//
func WithTopic(topic string) Option {
	return consumerOptionFunc(func(c *Consumer) {
		c.topic = topic
	})
}

//
func WithGroupID(id string) Option {
	return consumerOptionFunc(func(c *Consumer) {
		c.groupID = id
	})
}

const (
	AutoOffsetResetEarliest string = "earliest"
	AutoOffsetResetLatest   string = "latest"
	AutoOffsetResetNone     string = "none"
)

const DefaultAutoOffsetReset = AutoOffsetResetLatest

//
func WithAutoOffsetReset(reset string) Option {
	return consumerOptionFunc(func(c *Consumer) {
		c.autoOffsetReset = reset
	})
}

const (
	IsolationLevelReadUncommitted string = "read_uncommitted"
	IsolationLevelReadCommitted   string = "read_committed"
)

const DefaultIsolationLevel = IsolationLevelReadUncommitted

//
func WithIsolationLevel(lvl string) Option {
	return consumerOptionFunc(func(c *Consumer) {
		c.isolationLevel = lvl
	})
}

//
func WithClientID(id string) Option {
	return consumerOptionFunc(func(c *Consumer) {
		c.clientID = id
	})
}
