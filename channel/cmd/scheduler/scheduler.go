/*
Created By : Jannes Santoso
Scheduler Worker
Buffered Channel
Version 1.0
*/

package scheduler

import (
	"net/http"
	_ "net/http/pprof"
	"util/channel/cmd/scheduler/example/logic3"
)

func scheduler() {
	// logic1.RunScheduler()
	// logic2.RunScheduler()
	logic3.RunScheduler()

	http.ListenAndServe(":6060", nil)

}
