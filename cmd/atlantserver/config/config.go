package config

import (
	"github.com/caarlos0/env/v6"
)

//
type KafkaProductProducerConfig struct {
	//
	Servers []string `env:"KAFKA_PRODUCT_PRODUCER_SERVERS"`
}

//
type RPCServerConfig struct {
	//
	Host string `env:"RPC_SERVER_HOST"`
}

//
type MongoDBConfig struct {
	//
	URI string `env:"MONGODB_URI"`
}

//
type Config struct {
	KafkaProductProducerConfig
	RPCServerConfig
	MongoDBConfig
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

	if err = env.Parse(&c.MongoDBConfig); err != nil {
		return err
	}

	if err = env.Parse(c); err != nil {
		return err
	}

	return nil
}
