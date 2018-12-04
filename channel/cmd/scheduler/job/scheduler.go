package job

import "fmt"

const (
	stopped = iota
	running
)

var (
	mappingTasks = make(map[string]map[int]interface{})
)

var Action interface {
	Start()
	Pause()
}

func init() {
	Action = &scheduler{}
}

type scheduler struct {
	state     chan int
	routine   int
	nmRoutine string
	tasks     []interface{}
	input     chan interface{}
}

// Start - Start Scheduler
func (s *scheduler) Start() {
	(*s).state <- running
}

// Pause - Pause Scheduler
func (s *scheduler) Pause() {
	(*s).state <- stopped
}

func (s *scheduler) run() {
	state := (*s).state
	routine := (*s).routine
	nmRoutine := (*s).nmRoutine
	tasks := (*s).tasks
	input := (*s).input

	if len(mappingTasks[nmRoutine]) == 0 {
		mapTask := make(map[int]interface{})
		for k, v := range tasks {
			mapTask[k] = v
		}
		mappingTasks[nmRoutine] = mapTask
	}

	input = make(chan interface{})

	for i := 0; i < routine; i++ {
		go worker(input, state, nmRoutine)
	}

	go sendInput(mappingTasks[nmRoutine], input)
}

func monitoring(
	state chan int,
	routine int,
	nmRoutine string,
	input chan interface{},
) {

	for {
		select {
		case i := <-state:
			switch i {
			case running:
				fmt.Println("RUNNING ", nmRoutine, "TOTAL WORKER ", routine)

			case stopped:
				fmt.Println("STOPPED ", nmRoutine, "TOTAL WORKER ", routine)
				close(input)
				return
			}
		}
	}

}
