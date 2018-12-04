package job

import (
	"fmt"
)

// import "fmt"

type correlatedInputOutput struct {
	key  int
	data interface{}
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
		in := correlatedInputOutput{
			key:  k,
			data: v,
		}
		input <- in
	}
}

func getOutput(tasks map[int]interface{}, output chan int) {
	for i := 0; i < len(tasks); i++ {
		<-output
	}
}

// ChanInputData - Channel Receiver Data
type ChanInputData struct {
	State chan int
	Data  interface{}
}

func worker(
	input chan interface{},
	output chan int,
	state chan int,
	nmRoutine string,
) {
	for d := range input {
		data := d.(correlatedInputOutput)
		logicRun[nmRoutine].Run(ChanInputData{
			State: state,
			Data:  data.data,
		})

		output <- 1
	}
}
