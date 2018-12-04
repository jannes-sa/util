package job

import "fmt"

const (
	stopped = iota
	running
	done
)

var Action interface {
	Start(chan int)
	Pause(chan int)
}

func init() {
	Action = &scheduler{}
}

type scheduler struct{}

// Start - Start Scheduler
func (s scheduler) Start(state chan int) {
	state <- running
}

// Pause - Pause Scheduler
func (s scheduler) Pause(state chan int) {
	state <- stopped
}

func (s scheduler) run(
	state chan int,
	routine int,
	nmRoutine string,
	tasks []interface{},
) {

	input, output := make(chan interface{}), make(chan int)

	for i := 0; i < routine; i++ {
		go worker(input, output, state, nmRoutine)
	}

	mapTask := make(map[int]interface{})
	for k, v := range tasks {
		mapTask[k] = v
	}
	mappingTasks[nmRoutine] = mapTask

	go sendInput(mappingTasks[nmRoutine], input)
	getOutput(mappingTasks[nmRoutine], output)

	close(input)
	close(output)
}

func (s scheduler) monitoring(
	state chan int,
	routine int,
	nmRoutine string,
	tasks map[string]map[int]interface{},
	input chan interface{},
	output chan int,
) {

	for {
		select {
		case i := <-state:
			switch i {
			case running:
				fmt.Println("RUNNING ", nmRoutine, "TOTAL WORKER ", routine)

				input = make(chan interface{})

				go sendInput(mappingTasks[nmRoutine], input)
			case stopped:
				fmt.Println("STOPPED ", nmRoutine, "TOTAL WORKER ", routine)
				close(input)

			case done:
				fmt.Println("DONE ", nmRoutine, "TOTAL WORKER ", routine)
				close(input)
			}
		}
	}
}
