package main

import (
	"log"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("fileserver error: ", err)
	}

	logger = logger.With(zap.String("app", "fileserver"))

	rootCmd := cmdRoot(logger)

	if err = rootCmd.Execute(); err != nil {
		logger.Fatal("fileserver error: ", zap.Error(err))
	}
}
