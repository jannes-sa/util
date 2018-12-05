package scheduler

import (
	"time"
	"util/channel/lib/scheduler/example/logic1"
	"util/channel/lib/scheduler/example/logic2"
	"util/channel/lib/scheduler/example/logic3"
	"util/channel/lib/scheduler/example/logic4"
)

func scheduler() {
	logic1.RunScheduler()
	logic2.RunScheduler()
	logic3.RunScheduler()
	logic4.RunScheduler()

	time.Sleep(10 * time.Minute)
}
