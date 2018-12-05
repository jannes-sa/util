package job

type scheduler struct{}

func (s scheduler) run(
	routine int,
	nmRoutine string,
	tasks map[int]interface{},
	input chan interface{},
	output chan correlated,
) {
	mappingTasks[nmRoutine] = tasks

	input, output = make(chan interface{}), make(chan correlated)

	mappingStatusTasks[nmRoutine] = running
	for i := 0; i < routine; i++ {
		go worker(input, output, nmRoutine)
	}

	go sendinput(mappingTasks, nmRoutine, input)
	getOutput(len(mappingTasks[nmRoutine]), nmRoutine, output)

	close(input)
	close(output)

}
