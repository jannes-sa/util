package logic2

import (
	"fmt"
	"time"
	"util/channel/lib/scheduler/job"
)

const (
	logicNm string = "logic2"
)

func init() {
	job.RegisterLogic(logicNm, &logic2St{})
}

type logic2St struct{}

func (l logic2St) Run(receiverArg job.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))
}

func RunScheduler() {
	var tasks = []int{1, 2, 3, 4, 5, 6}

	var capsulateTasks []interface{}
	var t interface{}
	for _, v := range tasks {
		t = v
		capsulateTasks = append(capsulateTasks, t)
	}

	job.RunScheduler(1, logicNm, capsulateTasks)
}
