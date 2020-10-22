package main

import (
	"log"

	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

type CommandInitializer func(logger *zap.Logger) (cmd *cobra.Command)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("atlantclient error ", err)
	}

	logger = logger.With(zap.String("app", "atlantclient"))

	rootCmd := cmdRoot()

	for _, cmd := range []CommandInitializer{
		cmdFetch,
		cmdList,
	} {
		rootCmd.AddCommand(cmd(logger))
	}

	if err = rootCmd.Execute(); err != nil {
		logger.Error("atlantclient execute error: ", zap.Error(err))
	}
}

func cmdRoot() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "atlantclient",
	}

	return cmd
}
