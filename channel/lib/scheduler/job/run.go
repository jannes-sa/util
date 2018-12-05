package job

import (
	"fmt"
	"time"
)

var (
	mappingTasks       = make(map[string]map[int]interface{})
	mappingStatusTasks = make(map[string]status)
	debug              = true
)

type status uint8

const (
	preparing = iota
	running
	restart
	done
)

func (s status) String() string {
	str := map[status]string{
		preparing: "preparing",
		running:   "running",
		restart:   "restart",
		done:      "done",
	}

	v, ok := str[s]
	if !ok {
		return "UNDEFINED"
	}
	return v
}

// RunScheduler - Running Scheduler
func RunScheduler(
	worker int,
	nmWorker string,
	tasks []interface{},
) {
	var sch scheduler
	if !logicRun[nmWorker].Validate() {
		println("VALIDATE JOB", nmWorker, "FALSE")
		return
	}

	mappingStatusTasks[nmWorker] = preparing
	sch.run(worker, nmWorker, tasks)

	var mapWorker int
	go func() {
		for t := range time.Tick(5 * time.Second) {
			mapWorker = len(mappingTasks[nmWorker])
			print(t, nmWorker, "TOTAL TASKS LEFT", mapWorker, "STATUS", status.String(mappingStatusTasks[nmWorker]))

		}
	}()
}

func print(msg ...interface{}) {
	if debug {
		fmt.Println(msg)
	}
}
