package time

import (
	"time"
)

type Clock struct{}

func NewClock() (c *Clock) {
	return &Clock{}
}

func (c *Clock) NowInUTC() time.Time {
	return time.Now().UTC()
}
