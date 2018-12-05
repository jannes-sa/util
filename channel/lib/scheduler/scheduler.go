package scheduler

import (
	"time"
	"util/channel/lib/scheduler/example/logic4"
)

func scheduler() {
	logic4.RunScheduler()

	time.Sleep(10 * time.Minute)
}
