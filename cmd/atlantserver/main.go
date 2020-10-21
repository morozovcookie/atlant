package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/morozovcookie/atlant/cmd/atlantserver/config"
	"github.com/morozovcookie/atlant/grpc"
	svcV1 "github.com/morozovcookie/atlant/grpc/v1"
	"github.com/morozovcookie/atlant/http"
	"github.com/morozovcookie/atlant/http/client"
	"github.com/morozovcookie/atlant/kafka/producer"
	kafkaV1 "github.com/morozovcookie/atlant/kafka/v1"
	"github.com/morozovcookie/atlant/time"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	ggrpc "google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	logger = logger.With(zap.String("app", "atlantserver"))

	cfg := config.New()
	if err := cfg.Parse(); err != nil {
		logger.Fatal("parse config error", zap.Error(err))
	}

	c, err := initContainer(cfg, logger)
	if err != nil {
		logger.Fatal("init service container error", zap.Error(err))
	}

	atlantSvc := svcV1.NewAtlantService(
		c,
		logger.With(zap.String("component", "atlant_service")))

	s := grpc.NewServer(
		cfg.RPCServerConfig.Host,
		logger.With(zap.String("component", "grpc_server")),
		grpc.WithServiceRegistrator(func(gs *ggrpc.Server) {
			svcV1.RegisterAtlantServiceServer(gs, atlantSvc)
		}))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("starting application")

	eg := errgroup.Group{}
	eg.Go(s.Start)

	logger.Info("application started")

	<-quit

	logger.Info("stopping application")

	s.Stop()

	if err = eg.Wait(); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("application stopped")
}

func initContainer(cfg *config.Config, logger *zap.Logger) (c *svcV1.Container, err error) {
	c = &svcV1.Container{
		Clock: time.NewClock(),
	}

	c.ProductFetcher = http.NewProductFetcher(
		client.New(logger.With(zap.String("component", "http_client"))),
		logger.With(zap.String("component", "product_fetcher")))

	p, err := producer.New(
		logger.With(zap.String("component", "product_producer")),
		producer.WithServers(cfg.KafkaProductProducerConfig.Servers),
		producer.WithTopic(cfg.KafkaProductProducerConfig.Topic),
		producer.WithAcknowledgement(producer.WaitAllAcknowledgement),
		producer.WithTransactionalID("atlantserver"+cfg.Hostname),
		producer.WithIdempotenceState(producer.EnabledIdempotenceState))
	if err != nil {
		logger.Error("create producer error", zap.Error(err))

		return nil, err
	}

	c.ProductStorer = kafkaV1.NewProductStorer(
		p,
		logger.With(zap.String("component", "product_storer")))

	return c, nil
}
