package logic1

import (
	"fmt"
	"test/channel/cmd/ch1/scheduler"
	"time"
)

const (
	logicNm string = "logic1"
)

func init() {
	scheduler.RegisterLogic(logicNm, &logic1St{})
}

type logic1St struct{}

func (l logic1St) Run(receiverArg scheduler.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))

	if receiverArg.Data.(int) == 3 {
		scheduler.Action.Pause(receiverArg.State)
		time.Sleep(5 * time.Second)
		scheduler.Action.Start(receiverArg.State)
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
	scheduler.RunScheduler(c, 1, logicNm, capsulateTasks)
	scheduler.Action.Start(c)
}
