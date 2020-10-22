package mongodb

import (
	"github.com/pkg/errors"
)

var ErrTooMuchObjectsForStore = errors.New("could be store only one object at the same time")
