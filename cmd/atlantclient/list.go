package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func cmdList(_ *zap.Logger) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:     "list",
		Short:   "",
		Long:    "",
		Example: "atlantclient list --host 127.0.0.1:8080 --start 0 --limit 100 --sort ?",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
