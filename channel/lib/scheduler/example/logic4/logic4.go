package logic4

import (
	"errors"
	"fmt"
	"strconv"
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

func (l logic4St) Run(receiverArg job.ChanInputData) (
	resp interface{},
	err error,
) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(Tasks))
	if receiverArg.Data.(Tasks).task == 20 {
		err = errors.New("FAILED SOMETHING POKOKNYA")
		return
	}

	resp = "RESPONSE " + strconv.Itoa(receiverArg.Data.(Tasks).task)

	return
}

func (l logic4St) Done(out *job.OutputData) (state bool) {
	fmt.Println(
		"RESULT DONE", (*out).Result, "\n",
		"TOTAL TASK", (*out).TotalTasks, "\n",
		"TOTAL TASK DONE", (*out).TotalTasksDone, "\n",
		"TOTAL TASK FAIL", (*out).TotalTasksFail, "\n",
		"TOTAL TASK PENDING", (*out).TotalTasksPending, "\n",
	)

	for _, v := range (*out).Err {
		fmt.Println(
			"ERROR =>", v.Err, "\n",
			"INPUT TASK =>", v.InputError, "\n",
		)
	}

	return false
}

type Tasks struct {
	task       int
	taskString string
}

func RunScheduler() {

	var tasks []Tasks
	for i := 0; i < 100000; i++ {
		tasks = append(tasks, Tasks{task: i, taskString: "XXXXX"})
	}

	var capsulateTasks []interface{}
	var t interface{}
	for _, v := range tasks {
		t = v
		capsulateTasks = append(capsulateTasks, t)
	}

	err := job.RunScheduler(100, logicNm, capsulateTasks)
	if err != nil {
		fmt.Println(err)
	}
}
