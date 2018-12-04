/*
Created By : Jannes Santoso
Scheduler Worker
Buffered Channel
Version 1.0
*/
package scheduler

import (
	"fmt"
)

const (
	stopped = iota
	running
)

var (
	logicRun     = make(map[string]logiclayer)
	mappingTasks = make(map[string]map[int]interface{})
)

type logiclayer interface {
	Run(ChanInputData)
}

// RegisterLogic - Register Logic Inside Scheduler
func RegisterLogic(nmLogic string, logic logiclayer) {
	if _, ok := logicRun[nmLogic]; ok {
		msg := "failed Registered Logic " + nmLogic + "Already Registered"
		panic(msg)
		return
	}

	logicRun[nmLogic] = logic
}

var Action interface {
	Start(chan int)
	Pause(chan int)
}

func init() {
	Action = &Scheduler{}
}

type Scheduler struct{}

// RunScheduler - Running Scheduler
func RunScheduler(
	state chan int,
	routine int,
	nmRoutine string,
	tasks []interface{},
) {
	var sch Scheduler
	mapTask := make(map[int]interface{})
	for k, v := range tasks {
		mapTask[k] = v
	}
	mappingTasks[nmRoutine] = mapTask
	go sch.run(state, routine, nmRoutine, mappingTasks)
}

// Start - Start Scheduler
func (s Scheduler) Start(state chan int) {
	state <- running
}

// Pause - Pause Scheduler
func (s Scheduler) Pause(state chan int) {
	state <- stopped
}

func (s Scheduler) run(
	state chan int,
	routine int,
	nmRoutine string,
	tasks map[string]map[int]interface{},
) {
	var (
		input  chan interface{}
		output chan int
	)

	for {
		select {
		case i := <-state:
			switch i {
			case running:
				fmt.Println("RUNNING ", nmRoutine, "TOTAL WORKER ", routine)

				input, output = make(chan interface{}), make(chan int)

				for i := 0; i < routine; i++ {
					go worker(state, input, output, nmRoutine)
				}

				go sendinput(nmRoutine, input, tasks)
				// go getoutput(output)
			case stopped:
				fmt.Println("STOPPED ", nmRoutine, "TOTAL WORKER ", routine)
				close(input)
			}
		}
	}
}

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
