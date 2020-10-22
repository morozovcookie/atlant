package kafka

import (
	"github.com/pkg/errors"
)

var ErrDecodeIncomingMessage = errors.New("incoming message could not be decoded")
