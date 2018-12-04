package logic1

import (
	"fmt"
	"time"
	"util/channel/cmd/scheduler/job"
)

const (
	logicNm string = "logic1"
)

func init() {
	job.RegisterLogic(logicNm, &logic1St{})
}

type logic1St struct{}

func (l logic1St) Run(receiverArg job.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))

	if receiverArg.Data.(int) == 3 {
		// Pause job //
		job.Action.Pause(receiverArg.State)
		time.Sleep(5 * time.Second)
		// Continue job after 5 second Pause //
		job.Action.Start(receiverArg.State)
	}

}

func RunScheduler() {
	var tasks = []int{1, 2, 3, 4, 5, 6}

	var capsulateTasks []interface{}
	var t interface{}
	for _, v := range tasks {
		t = v
		capsulateTasks = append(capsulateTasks, t)
	}

	c := make(chan int)
	job.RunScheduler(c, 5, logicNm, capsulateTasks)
	job.Action.Start(c)
}
