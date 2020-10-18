package cmd

import (
	"context"
	stdhttp "net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/morozovcookie/atlant/http"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	ErrEmptyHostFlagPath = errors.New("empty host flag")
	ErrInvalidHostValue  = errors.New("invalid host value")
	ErrEmptyDirPathFlag  = errors.New("empty dirpath flag")
	ErrFileDoesNotExist  = errors.New("file does not exist")
	ErrNotDir            = errors.New("specified file is not a directory")
)

const (
	CancelTimeout = time.Second * 5

	FileServerHandlerPattern = "/"
)

type Flag interface {
	ValidateFlag() (err error)
}

type HostFlag string

func (f *HostFlag) ValidateFlag() (err error) {
	if f.String() == "" {
		return ErrEmptyHostFlagPath
	}

	r, err := regexp.Compile(`^(((\d{1,3}\.){3}(\d{1,3}))|((\w+\.){2}(\w+))):(\d{4,5})$`)
	if err != nil {
		return err
	}

	if !r.MatchString(f.String()) {
		return errors.New(ErrInvalidHostValue.Error() + ": " + f.String())
	}

	return nil
}

func (f *HostFlag) String() (s string) {
	return (string)(*f)
}

func (f *HostFlag) Pointer() (p *string) {
	return (*string)(f)
}

type DirPathFlag string

func (f *DirPathFlag) ValidateFlag() (err error) {
	if f.String() == "" {
		return ErrEmptyDirPathFlag
	}

	i, err := os.Stat(f.String())
	if os.IsNotExist(err) {
		return ErrFileDoesNotExist
	}

	if !i.IsDir() {
		return ErrNotDir
	}

	return nil
}

func (f *DirPathFlag) String() (s string) {
	return (string)(*f)
}

func (f *DirPathFlag) Pointer() (p *string) {
	return (*string)(f)
}

type RootCommandOptions struct {
	host    *HostFlag
	dirPath *DirPathFlag

	logger *zap.Logger
}

func (opts *RootCommandOptions) Validate() (err error) {
	for _, f := range []Flag{opts.host, opts.dirPath} {
		if err = f.ValidateFlag(); err != nil {
			return err
		}
	}

	return nil
}

func (opts *RootCommandOptions) Run() (err error) {
	s := http.NewServer(opts.host.String(),
		http.WithHandler(FileServerHandlerPattern, stdhttp.FileServer(stdhttp.Dir(opts.dirPath.String()))))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	opts.logger.Debug("starting file server")

	eg := errgroup.Group{}
	eg.Go(s.Start)

	<-quit

	opts.logger.Debug("stopping file server")

	stopCtx, cancel := context.WithTimeout(context.Background(), CancelTimeout)
	defer cancel()

	if err = s.Stop(stopCtx); err != nil {
		return err
	}

	return eg.Wait()
}

func NewRootCommand(logger *zap.Logger) (c *cobra.Command) {
	opts := &RootCommandOptions{
		host:    new(HostFlag),
		dirPath: new(DirPathFlag),

		logger: logger,
	}

	c = &cobra.Command{
		Use:     "fileserver",
		Short:   "Fileserver is just a file server",
		Long:    "Fileserver give you opportunity to serve files from your specific directory",
		Example: "fileserver --host 0.0.0.0 --dirpath /opt",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Validate(); err != nil {
				return err
			}

			return opts.Run()
		},
	}

	c.Flags().StringVar(opts.host.Pointer(), "host", "", "file server host")
	c.Flags().StringVar(opts.dirPath.Pointer(), "dirpath", "", "file server directory")

	return c
}
