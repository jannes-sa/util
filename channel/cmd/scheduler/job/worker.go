package job

type correlatedInput struct {
	key int
	val interface{}
}

func sendInput(tasks map[int]interface{}, input chan interface{}) {
	go func() {
		for k, v := range tasks {
			in := correlatedInput{
				key: k,
				val: v,
			}
			input <- in
		}
	}()
}

func getOutput(tasks map[int]interface{}, output chan int) {
	for i := 0; i < len(tasks); i++ {
		<-output
	}
}

func worker(
	input chan interface{},
	output chan int,
	state chan int,
	nmRoutine string,
) {
	for data := range input {
		d := data.(correlatedInput)
		logicRun[nmRoutine].Run(ChanInputData{
			State: state,
			Data:  d.val,
		})

		output <- d.key
	}
}
