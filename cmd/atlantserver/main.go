package main

import (
	"context"
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
	"github.com/morozovcookie/atlant/mongodb"
	"github.com/morozovcookie/atlant/time"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	ggrpc "google.golang.org/grpc"
)

const (
	appname string = "atlantserver"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	logger = logger.With(zap.String("app", appname))

	cfg := config.New()
	if err := cfg.Parse(); err != nil {
		logger.Fatal("parse config error", zap.Error(err))
	}

	p, err := initProducer(cfg, logger)
	if err != nil {
		logger.Fatal("create producer error", zap.Error(err))
	}

	mc, err := initMongoDB(cfg)
	if err != nil {
		logger.Fatal("create mongodb client error", zap.Error(err))
	}

	var s *grpc.Server
	{
		atlantSvc := svcV1.NewAtlantService(
			initContainer(p, mc, logger),
			logger.With(zap.String("component", "atlant_service")))

		s = grpc.NewServer(
			cfg.RPCServerConfig.Host,
			logger.With(zap.String("component", "grpc_server")),
			grpc.WithServiceRegistrator(func(gs *ggrpc.Server) {
				svcV1.RegisterAtlantServiceServer(gs, atlantSvc)
			}))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("starting application")

	eg := errgroup.Group{}
	eg.Go(s.Start)

	logger.Info("application started")

	<-quit

	logger.Info("stopping application")

	p.Close(context.Background())

	s.Stop()

	if err = mc.Close(context.Background()); err != nil {
		logger.Error("mongodb client close error", zap.Error(err))
	}

	if err = eg.Wait(); err != nil {
		logger.Error("stopping application error", zap.Error(err))
	}

	logger.Info("application stopped")
}

func initContainer(p kafkaV1.Producer, mc mongodb.MongoCollector, logger *zap.Logger) (c *svcV1.Container) {
	c = &svcV1.Container{
		Clock: time.NewClock(),
	}

	c.ProductFetcher = http.NewProductFetcher(
		client.New(logger.With(zap.String("component", "http_client"))),
		logger.With(zap.String("component", "product_fetcher")))

	c.ProductStorer = kafkaV1.NewProductStorer(
		p,
		logger.With(zap.String("component", "product_storer")))

	c.ProductLister = mongodb.NewProductStorage(
		mc,
		logger.With(zap.String("component", "product_lister")))

	return c
}

func initProducer(cfg *config.Config, logger *zap.Logger) (p *producer.Producer, err error) {
	p, err = producer.New(
		logger.With(zap.String("component", "product_producer")),
		producer.WithServers(cfg.KafkaProductProducerConfig.Servers),
		producer.WithTopic("docker.atlant.cdc.products.0"),
		producer.WithAcknowledgement(producer.AcknowledgementWaitAll),
		producer.WithTransactionalID(appname+"-"+cfg.Hostname),
		producer.WithIdempotenceState(producer.IdempotenceEnabledState),
		producer.WithCompressionType(producer.CompressionTypeGzip))
	if err != nil {
		return nil, err
	}

	if err = p.InitTransactions(context.Background()); err != nil {
		p.Close(context.Background())

		return nil, err
	}

	return p, nil
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
