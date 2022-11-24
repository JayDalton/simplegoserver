package metrics

import "time"

var startTime time.Time

func init() {
	startTime = time.Now()
}

func Uptime() (int64, string) {
	return time.Since(startTime).Milliseconds(), Milliseconds
}
