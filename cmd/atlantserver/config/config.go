package config

import (
	"github.com/caarlos0/env/v6"
)

//
type KafkaProductProducerConfig struct {
	//
	Servers []string `env:"KAFKA_PRODUCT_PRODUCER_SERVERS"`

	//
	Topic string `env:"KAFKA_PRODUCT_PRODUCER_TOPIC"`
}

//
type RPCServerConfig struct {
	//
	Host string `env:"RPC_SERVER_HOST"`
}

//
type Config struct {
	Hostname string `env:"HOSTNAME"`

	KafkaProductProducerConfig
	RPCServerConfig
}

//
func New() *Config {
	return &Config{}
}

//
func (c *Config) Parse() (err error) {
	if err = env.Parse(&c.KafkaProductProducerConfig); err != nil {
		return err
	}

	if err = env.Parse(&c.RPCServerConfig); err != nil {
		return err
	}

	if err = env.Parse(c); err != nil {
		return err
	}

	return nil
}
