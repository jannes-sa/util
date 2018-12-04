package logic4

import (
	"fmt"
	"time"
	"util/channel/cmd/scheduler/job"
)

const (
	logicNm string = "logic4"
)

func init() {
	job.RegisterLogic(logicNm, &logic4St{})
}

type logic4St struct{}

var ck []int

func (l logic4St) Run(receiverArg job.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))
	ck = append(ck, 1)
	if receiverArg.Data.(int) == 5 {
		job.Action.Pause()
		time.Sleep(10 * time.Second)
		job.Action.Start()
	}

}

func RunScheduler() {

	var tasks []int
	for i := 0; i < 10; i++ {
		tasks = append(tasks, i)
	}

	var capsulateTasks []interface{}
	var t interface{}
	for _, v := range tasks {
		t = v
		capsulateTasks = append(capsulateTasks, t)
	}

	c := make(chan int)
	job.RunScheduler(c, 1, logicNm, capsulateTasks)

	fmt.Println("DONE TOTAL", len(ck))
}
