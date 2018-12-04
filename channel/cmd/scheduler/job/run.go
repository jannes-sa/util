package job

// RunScheduler - Running Scheduler
func RunScheduler(
	state chan int,
	worker int,
	nmWorker string,
	tasks []interface{},
) {
	input := make(chan interface{})

	sch := scheduler{
		state:     state,
		routine:   worker,
		nmRoutine: nmWorker,
		tasks:     tasks,
		input:     input,
	}

	sch.run()

}
