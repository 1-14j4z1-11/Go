package performance

import (
	"time"
)

type Action func()

func MeasurePerformance(action Action) time.Duration {
	start := time.Now()
	action()
	end := time.Now()

	return end.Sub(start)
}
