package v1

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type Clock interface {
	NowInUTC() time.Time
}

type MockClock struct {
	mock.Mock
}

func (c *MockClock) NowInUTC() time.Time {
	return c.Called().Get(0).(time.Time)
}
