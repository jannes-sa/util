package job

import (
	"fmt"
	"time"
)

// func sendinput(
// 	nmRoutine string,
// 	input chan interface{},
// 	tasks map[string]map[int]interface{},
// ) {
// 	// defer func() {
// 	// 	// recover from panic caused by writing to a closed channel
// 	// 	if r := recover(); r != nil {
// 	// 		err := fmt.Errorf("%v", r)
// 	// 		fmt.Printf("write: error writing %d on channel: %v\n", input, err)
// 	// 		return
// 	// 	}
// 	// }()

// 	for k, v := range tasks[nmRoutine] {
// 		input <- v
// 		delete(tasks[nmRoutine], k)
// 	}
// }

// func getOutput(count int, output chan int) {
// 	for i := 0; count < 0; i++ {
// 		<-output
// 	}
// }

// func worker(
// 	input chan interface{},
// 	output chan int,
// 	state chan int,
// 	nmRoutine string,
// ) {
// 	for data := range input {
// 		logicRun[nmRoutine].Run(ChanInputData{
// 			State: state,
// 			Data:  data,
// 		})
// 		output <- 1
// 	}
// }

func sendInput(tasks map[int]interface{}, input chan interface{}) {
	go func() {
		for _, v := range tasks {
			input <- v
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
		fmt.Println(data, time.Now())

		output <- data.(int)
	}
}
