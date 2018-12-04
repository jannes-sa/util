package logic3

import (
	"fmt"
	"time"
	"util/channel/cmd/scheduler/job"
)

const (
	logicNm string = "logic3"
)

func init() {
	job.RegisterLogic(logicNm, &logic3St{})
}

type logic3St struct{}

func (l logic3St) Run(receiverArg job.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))

	if receiverArg.Data.(int) == 5 {
		job.Action.Pause(receiverArg.State)
		time.Sleep(10 * time.Second)
		job.Action.Start(receiverArg.State)
	}

}

func RunScheduler() {
	var tasks []int
	for i := 0; i <= 100; i++ {
		tasks = append(tasks, i)
	}

	var capsulateTasks []interface{}
	var t interface{}
	for _, v := range tasks {
		t = v
		capsulateTasks = append(capsulateTasks, t)
	}

	c := make(chan int)
	job.RunScheduler(c, 2, logicNm, capsulateTasks)
	job.Action.Start(c)
}
