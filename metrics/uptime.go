package metrics

import "time"

var startTime time.Time

func init() {
	startTime = time.Now()
}

func Uptime() time.Duration {
	return time.Since(startTime)
}
