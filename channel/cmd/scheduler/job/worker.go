package job

import "fmt"

func sendinput(
	nmRoutine string,
	input chan interface{},
	tasks map[string]map[int]interface{},
) {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			fmt.Printf("write: error writing %d on channel: %v\n", input, err)
			return
		}
	}()

	for k, v := range tasks[nmRoutine] {
		input <- v
		delete(tasks[nmRoutine], k)
	}

}

// func getoutput(output chan int) {
// 	for {
// 		fmt.Println(<-output)
// 	}
// }

// ChanInputData - Channel Receiver Data
type ChanInputData struct {
	State chan int
	Data  interface{}
}

func worker(
	state chan int,
	input chan interface{},
	output chan int,
	nmRoutine string,
) {
	for data := range input {

		logicRun[nmRoutine].Run(ChanInputData{
			State: state,
			Data:  data,
		})
		// output <- data
	}
}
