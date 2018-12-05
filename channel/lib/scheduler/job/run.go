package job

import (
	"fmt"
	"time"
)

var (
	mappingTasks = make(map[string]map[int]interface{})
	debug        = true
)

// RunScheduler - Running Scheduler
func RunScheduler(
	worker int,
	nmWorker string,
	tasks []interface{},
) {
	var sch scheduler
	sch.run(worker, nmWorker, tasks)

	var mapWorker int
	go func() {
		for t := range time.Tick(5 * time.Second) {
			mapWorker = len(mappingTasks[nmWorker])
			print(t, nmWorker, "TOTAL TASKS LEFT", mapWorker)
		}
	}()
}

func print(msg ...interface{}) {
	if debug {
		fmt.Println(msg)
	}
}
