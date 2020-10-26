package main

import (
	"context"
	stderrors "errors"
	"fmt"
	"io"
	"os"
	"strings"

	v1 "github.com/morozovcookie/atlant/grpc/v1"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type ListCommandOptions struct {
	host *HostFlag

	start *StartFlag

	limit *LimitFlag

	sort *SortFlag

	crt string
}

func (opts *ListCommandOptions) Validate() (err error) {
	for _, opt := range []Flag{opts.host, opts.start, opts.limit, opts.sort} {
		err = opt.Validate()

		if stderrors.Is(ErrInvalidLimitParameterMinValue, err) && opts.limit.Int64() == 0 {
			opts.limit = new(LimitFlag)
			*(opts.limit.Pointer()) = 100

			continue
		}

		if stderrors.Is(ErrInvalidLimitParameterMaxValue, err) {
			opts.limit = new(LimitFlag)
			*(opts.limit.Pointer()) = 100

			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (opts ListCommandOptions) Run(ctx context.Context, logger *zap.Logger) (err error) {
	dialOpt := grpc.WithInsecure()

	if opts.crt != "" {
		creds, err := credentials.NewServerTLSFromFile(opts.crt, "")
		if err != nil {
			return err
		}

		dialOpt = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.DialContext(ctx, opts.host.String(), dialOpt)
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

	pp, err := opts.makeListRequest(ctx, svc, logger)

	for _, p := range pp {
		_, _ = fmt.Fprintf(os.Stdout, "%+v \n", p)
	}

	return nil
}

func (opts ListCommandOptions) makeListRequest(
	ctx context.Context,
	svc v1.AtlantServiceClient,
	logger *zap.Logger,
) (
	pp []*v1.ListResponse_Product,
	err error,
) {
	req := &v1.ListRequest{
		Start:   opts.start.Int64(),
		Limit:   opts.limit.Int64(),
		Options: make([]*v1.ListRequest_SortingOption, len(opts.sort.StringArray())),
	}

	for i, opt := range opts.sort.StringArray() {
		var (
			ss = strings.Split(opt, ":")

			field     = ss[0]
			direction = ss[1]
		)

		req.Options[i] = &v1.ListRequest_SortingOption{
			Field:     field,
			Direction: 1,
		}

		if direction == "desc" {
			req.Options[i].Direction = 2
		}
	}

	res, err := svc.List(ctx, req)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			logger.Error("request error",
				zap.NamedError("error", grpcErr.Err()),
				zap.String("code", grpcErr.Code().String()),
				zap.String("message", grpcErr.Message()),
				zap.Any("details", grpcErr.Details()))
		}

		return nil, err
	}

	return res.Products, nil
}

func cmdList(logger *zap.Logger) (cmd *cobra.Command) {
	opts := &ListCommandOptions{
		host:  new(HostFlag),
		start: new(StartFlag),
		limit: new(LimitFlag),
		sort:  new(SortFlag),
	}

	cmd = &cobra.Command{
		Use:     "list",
		Short:   "",
		Long:    "",
		Example: "atlantclient list --host 127.0.0.1:8080 --start 0 --limit 100 --sort name:desc,updated_at:asc",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if err = opts.Validate(); err != nil {
				return err
			}

			return opts.Run(cmd.Context(), logger)
		},
	}

	cmd.Flags().StringVar(opts.host.Pointer(), "host", "", "server host")
	cmd.Flags().Int64Var(opts.start.Pointer(), "start", 0, "start position")
	cmd.Flags().Int64Var(opts.limit.Pointer(), "limit", 0, "items per page")
	cmd.Flags().StringSliceVar(opts.sort.Pointer(), "sort", nil, "sorting parameters")
	cmd.Flags().StringVar(&opts.crt, "crt", "", "server certificate")

	return cmd
}
