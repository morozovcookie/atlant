package config

import (
	"github.com/caarlos0/env"
)

//
type KafkaProductConsumerConfig struct {
	//
	Servers []string `env:"KAFKA_PRODUCT_CONSUMER_SERVERS"`
}

//
type MongoDBConfig struct {
	//
	URI string `env:"MONGODB_URI"`
}

//
type Config struct {
	KafkaProductConsumerConfig
	MongoDBConfig
}

//
func New() (c *Config) {
	return &Config{}
}

//
func (c *Config) Parse() (err error) {
	if err = env.Parse(&c.KafkaProductConsumerConfig); err != nil {
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
