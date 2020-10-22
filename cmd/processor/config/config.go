package config

import (
	"github.com/caarlos0/env"
)

//
type KafkaProductProducerConfig struct {
	//
	Servers []string `env:"KAFKA_PRODUCT_PRODUCER_SERVERS"`
}

//
type MongoDBConfig struct {
	URI string `env:"MONGODB_URI"`
}

//
type Config struct {
	KafkaProductProducerConfig
	MongoDBConfig
}

func New() (c *Config) {
	return &Config{}
}

func (c *Config) Parse() (err error) {
	if err = env.Parse(&c.KafkaProductProducerConfig); err != nil {
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
