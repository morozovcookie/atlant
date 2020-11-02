package main

import (
	"context"
	"io"

	v1 "github.com/morozovcookie/atlant/grpc/v1"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type FetchCommandOptions struct {
	host *HostFlag

	url *URLFlag

	crt string
}

func (opts *FetchCommandOptions) Validate() (err error) {
	for _, opt := range []Flag{opts.host, opts.url} {
		if err = opt.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (opts *FetchCommandOptions) Run(ctx context.Context, logger *zap.Logger) (err error) {
	conn, err := grpc.DialContext(ctx, opts.host.String(), grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer func(closer io.Closer) {
		if closeErr := closer.Close(); closeErr != nil {
			logger.Error("close conn error ", zap.Error(closeErr))
			err = closeErr
		}
	}(conn)

	svc := v1.NewAtlantServiceClient(conn)

	_, err = svc.Fetch(
		ctx,
		&v1.FetchRequest{
			Url: opts.url.String(),
		})
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			logger.Error("request error",
				zap.NamedError("error", grpcErr.Err()),
				zap.String("code", grpcErr.Code().String()),
				zap.String("message", grpcErr.Message()),
				zap.Any("details", grpcErr.Details()))
		}

		return err
	}

	return nil
}

func cmdFetch(logger *zap.Logger) (cmd *cobra.Command) {
	opts := &FetchCommandOptions{
		host: new(HostFlag),
		url:  new(URLFlag),
	}

	cmd = &cobra.Command{
		Use:     "fetch",
		Short:   "",
		Long:    "",
		Example: "atlantclient fetch --host 127.0.0.1:8080 --url http://example.com/sample.csv",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if err = opts.Validate(); err != nil {
				return err
			}

			return opts.Run(cmd.Context(), logger)
		},
	}

	cmd.Flags().StringVar(opts.host.Pointer(), "host", "", "server host")
	cmd.Flags().StringVar(opts.url.Pointer(), "url", "", "csv resource url")
	cmd.Flags().StringVar(&opts.crt, "crt", "", "server certificate")

	return cmd
}
