package job

import (
	"errors"
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
	logic logiclayer,
) (err error) {
	err = registerLogic(nmWorker, logic)
	if err != nil {
		return
	}

	mappingStatusTasks[nmWorker] = preparing
	err = prepareRun(worker, nmWorker, tasks)
	if err != nil {
		return
	}

	var mapWorker int
	go func() {
		for t := range time.Tick(5 * time.Second) {
			mapWorker = len(mappingTasks[nmWorker])
			print(t, nmWorker, "TOTAL TASKS LEFT", mapWorker, "STATUS", status.String(mappingStatusTasks[nmWorker]))
			if mappingStatusTasks[nmWorker] == restart {
				err = prepareRun(worker, nmWorker, tasks)
				if err != nil {
					return
				}
			}
		}
	}()

	return
}

func prepareRun(
	worker int,
	nmWorker string,
	tasks []interface{},
) (err error) {
	var sch scheduler
	if !logicRun[nmWorker].Validate() {
		println("VALIDATE JOB", nmWorker, "FALSE")
		err = errors.New("VALIDATE JOB" + nmWorker + "FALSE")
		return
	}
	sch.run(worker, nmWorker, tasks)
	return
}

func print(msg ...interface{}) {
	if debug {
		fmt.Println(msg)
	}
}
