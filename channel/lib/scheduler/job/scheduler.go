package job

type scheduler struct{}

func (s scheduler) run(
	routine int,
	nmRoutine string,
	tasks []interface{},
) {
	mapTask := make(map[int]interface{})
	for k, v := range tasks {
		mapTask[k] = v
	}
	mappingTasks[nmRoutine] = mapTask

	input, output := make(chan interface{}), make(chan int)

	mappingStatusTasks[nmRoutine] = running
	for i := 0; i < routine; i++ {
		go worker(input, output, nmRoutine)
	}

	go sendinput(mappingTasks, nmRoutine, input)
	getOutput(len(mappingTasks[nmRoutine]), nmRoutine, output)

	close(input)
	close(output)

}
