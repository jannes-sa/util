/*
Created By : Jannes Santoso
Scheduler Worker
Buffered Channel
Version 1.0
*/

package scheduler

import (
	"time"
	"util/channel/cmd/scheduler/logic3"
)

func scheduler() {
	// logic1.RunScheduler()
	// logic2.RunScheduler()
	logic3.RunScheduler()

	time.Sleep(10 * time.Minute)
}
