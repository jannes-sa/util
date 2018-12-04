package job

// RunScheduler - Running Scheduler
func RunScheduler(
	state chan int,
	worker int,
	nmWorker string,
	tasks []interface{},
) {
	var input chan interface{}

	sch := scheduler{
		state:     state,
		routine:   worker,
		nmRoutine: nmWorker,
		tasks:     tasks,
		input:     input,
	}

	monitoring(state, worker, nmWorker, input)
	sch.Start()

}
