package job

var (
	mappingTasks = make(map[string]map[int]interface{})
)

// RunScheduler - Running Scheduler
func RunScheduler(
	state chan int,
	worker int,
	nmWorker string,
	tasks []interface{},
) {
	var sch scheduler

	go sch.run(state, worker, nmWorker, tasks)
}
