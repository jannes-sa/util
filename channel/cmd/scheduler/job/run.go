package job

var (
	mappingTasks = make(map[string]map[int]interface{})
)

// RunScheduler - Running Scheduler
func RunScheduler(
	state chan int,
	routine int,
	nmRoutine string,
	tasks []interface{},
) {
	var (
		sch scheduler
	)

	sch.run(
		state,
		routine,
		nmRoutine,
		tasks,
	)

}
