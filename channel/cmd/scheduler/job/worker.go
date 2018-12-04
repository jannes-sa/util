package job

import "fmt"

type correlatedInput struct {
	key int
	val interface{}
}

func sendInput(tasks map[int]interface{}, input chan interface{}) {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			fmt.Printf("write: error writing %d on channel: %v\n", input, err)
			return
		}
	}()

	for k, v := range tasks {
		in := correlatedInput{
			key: k,
			val: v,
		}
		input <- in
	}

}

func getOutput(tasks map[int]interface{}, nmRoutine string, output chan int) {
	for i := 0; i < len(tasks); i++ {
		<-output
		// delete(mappingTasks[nmRoutine], o)
	}
}

func worker(
	input chan interface{},
	state chan int,
	nmRoutine string,
) {
	for data := range input {
		d := data.(correlatedInput)

		logicRun[nmRoutine].Run(ChanInputData{
			State: state,
			Data:  d.val,
		})
	}
}
