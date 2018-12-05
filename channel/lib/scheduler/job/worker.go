package job

import "fmt"

type correlatedInput struct {
	key int
	val interface{}
}

func sendinput(
	mappingTasks map[string]map[int]interface{},
	nmRoutine string,
	input chan interface{},
) {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			fmt.Printf("write: error writing %d on channel: %v\n", input, err)
			return
		}
	}()

	for k, v := range mappingTasks[nmRoutine] {
		in := correlatedInput{
			key: k,
			val: v,
		}
		input <- in
		delete(mappingTasks[nmRoutine], k)
	}
}

func getOutput(countTasks int, nmRoutine string, output chan int) {
	var workDone []int
	for i := 0; i < countTasks; i++ {
		o := <-output
		workDone = append(workDone, o)
	}
	println(nmRoutine, "JOB DONE FROM ", countTasks, " = ", len(workDone))

	var out OutputData
	out.TotalTasks = len(workDone)
	out.Result = workDone

	if !logicRun[nmRoutine].Done(&out) {
		mappingStatusTasks[nmRoutine] = restart
		return
	}
	mappingStatusTasks[nmRoutine] = done
}

// ChanInputData - Channel Receiver Data
type ChanInputData struct {
	Data interface{}
}

type OutputData struct {
	TotalTasks int
	Result     interface{}
}

func worker(
	input chan interface{},
	output chan int,
	nmRoutine string,
) {
	for data := range input {
		d := data.(correlatedInput)
		logicRun[nmRoutine].Run(ChanInputData{
			Data: d.val,
		})
		output <- d.key
	}
}
