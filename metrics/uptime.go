package metrics

import "time"

type Uptime struct {
	start time.Time
}

func NewUptime(time time.Time) *Uptime {
	return &Uptime{
		start: time,
	}
}

func (uptime *Uptime) Value() (int64, string) {
	return time.Since(uptime.start).Milliseconds(), Milliseconds
}
