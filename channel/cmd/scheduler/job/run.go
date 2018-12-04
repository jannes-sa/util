package job

var (
	mappingTasks = make(map[string]map[int]interface{})
)

// RunScheduler - Running Scheduler
func RunScheduler(
	state chan int,
	routine int,
	nmRoutine string,
	tasks []interface{},
) {
	var (
		// 	input  chan interface{}
		// 	output chan int
		sch scheduler
	)

	sch.run(
		state,
		routine,
		nmRoutine,
		tasks,
	)

}

// func sendInput(tasks map[int]interface{}, input chan interface{}) {
// 	for _, v := range tasks {
// 		input <- v.(int)
// 	}
// }
