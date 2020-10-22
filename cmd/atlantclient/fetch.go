package main

import (
	"context"
	"io"
	"regexp"

	v1 "github.com/morozovcookie/atlant/grpc/v1"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	ErrEmptyHostFlag    = errors.New("empty host flag")
	ErrInvalidHostValue = errors.New("invalid host value")
	ErrEmptyURLFlag     = errors.New("empty URL flag")
	ErrInvalidURLValue  = errors.New("invalid URL value")
)

type Flag interface {
	Validate() (err error)
}

type FetchCommandHostFlag string

func (f FetchCommandHostFlag) String() (s string) {
	return (string)(f)
}

func (f *FetchCommandHostFlag) Pointer() (p *string) {
	return (*string)(f)
}

func (f FetchCommandHostFlag) Validate() (err error) {
	if f.String() == "" {
		return ErrEmptyHostFlag
	}

	r := regexp.MustCompile(`^(((\d{1,3}\.){3}(\d{1,3}))|((\w+\.){2}(\w+))):(\d{4,5})$`)

	if !r.MatchString(f.String()) {
		return errors.New(ErrInvalidHostValue.Error() + ": " + f.String())
	}

	return nil
}

type FetchCommandURLFlag string

func (f FetchCommandURLFlag) String() (s string) {
	return (string)(f)
}

func (f *FetchCommandURLFlag) Pointer() (p *string) {
	return (*string)(f)
}

func (f FetchCommandURLFlag) Validate() (err error) {
	if f.String() == "" {
		return ErrEmptyURLFlag
	}

	r := regexp.MustCompile(`(http|https)://[\w\-_]+(\.[\w\-_]+)+([\w\-.,@?^=%&amp;:/~+#]*[\w\-@?^=%&amp;/~+#])?`)

	if !r.MatchString(f.String()) {
		return errors.New(ErrInvalidURLValue.Error() + ": " + f.String())
	}

	return nil
}

type FetchCommandOptions struct {
	host *FetchCommandHostFlag

	url *FetchCommandURLFlag
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
			URL: opts.url.String(),
		})
	if err != nil {
		return err
	}

	return nil
}

func cmdFetch(logger *zap.Logger) (cmd *cobra.Command) {
	opts := &FetchCommandOptions{
		host: new(FetchCommandHostFlag),
		url:  new(FetchCommandURLFlag),
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

	return cmd
}
