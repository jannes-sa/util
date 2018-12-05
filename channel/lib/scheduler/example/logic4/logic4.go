package logic4

import (
	"fmt"
	"time"
	"util/channel/lib/scheduler/job"
)

const (
	logicNm string = "logic4"
)

func init() {
	job.RegisterLogic(logicNm, &logic4St{})
}

type logic4St struct{}

func (l logic4St) Validate() (state bool) {
	return true
}

func (l logic4St) Run(receiverArg job.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))
}

func (l logic4St) Done(out *job.OutputData) (state bool) {
	fmt.Println("RESULT FROM DONE", (*out).TotalTasks)
	return true
}

func RunScheduler() {

	var tasks []int
	for i := 0; i < 1000; i++ {
		tasks = append(tasks, i)
	}

	var capsulateTasks []interface{}
	var t interface{}
	for _, v := range tasks {
		t = v
		capsulateTasks = append(capsulateTasks, t)
	}

	job.RunScheduler(100, logicNm, capsulateTasks)
}
