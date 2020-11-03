package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/morozovcookie/atlant/cmd/processor/config"
	"github.com/morozovcookie/atlant/kafka/consumer"
	v1 "github.com/morozovcookie/atlant/kafka/v1"
	"github.com/morozovcookie/atlant/mongodb"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	appname string = "product-processor"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	logger = logger.With(zap.String("app", appname))

	cfg := config.New()

	if err = cfg.Parse(); err != nil {
		logger.Fatal("config parse error", zap.Error(err))
	}

	c, err := initConsumer(cfg, logger)
	if err != nil {
		logger.Fatal("init consumer error", zap.Error(err))
	}

	mc, err := initMongoDB(cfg)
	if err != nil {
		logger.Fatal("init mongodb client error", zap.Error(err))
	}

	pp := v1.NewProductProcessor(
		mongodb.NewProductStorage(mc, logger.With(zap.String("component", "product_storage"))),
		logger.With(zap.String("component", "product_processor")))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("starting application")

	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return c.Subscribe(pp.ProcessProduct)
	})

	logger.Info("application started")

	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}

	logger.Info("application stopping")

	if err = c.Close(); err != nil {
		logger.Error("closing consumer error", zap.Error(err))
	}

	if err = mc.Close(context.Background()); err != nil {
		logger.Error("mongodb close error", zap.Error(err))
	}

	if err = eg.Wait(); err != nil {
		logger.Error("stopping application error", zap.Error(err))
	}

	logger.Info("application stopped")
}

func initConsumer(cfg *config.Config, logger *zap.Logger) (c *consumer.Consumer, err error) {
	return consumer.New(
		context.Background(),
		logger.With(zap.String("component", "product_consumer")),
		consumer.WithServers(cfg.KafkaProductConsumerConfig.Servers),
		consumer.WithTopic("docker.atlant.cdc.products.0"),
		consumer.WithGroupID(appname),
		consumer.WithIsolationLevel(consumer.IsolationLevelReadCommitted),
		consumer.WithAutoOffsetReset(consumer.AutoOffsetResetEarliest))
}

func initMongoDB(cfg *config.Config) (mc *mongodb.Client, err error) {
	mc = mongodb.NewClient(
		mongodb.WithURI(cfg.MongoDBConfig.URI),
		mongodb.WithDatabase("atlant"),
		mongodb.WithCollection("products"))

	if err = mc.Connect(context.Background()); err != nil {
		return nil, err
	}

	return mc, nil
}
