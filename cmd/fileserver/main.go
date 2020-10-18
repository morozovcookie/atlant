package main

import (
	"fmt"
	"os"

	"github.com/morozovcookie/atlant/cmd/fileserver/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fileserver: %v", err)
	}

	logger = logger.With(zap.String("app", "fileserver"))

	if err := cmd.NewRootCommand(logger).Execute(); err != nil {
		logger.Fatal("fileserver error: ", zap.Error(err))
	}
}
