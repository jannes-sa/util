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

func (l logic4St) Run(receiverArg job.ChanInputData) {
	fmt.Println(time.Now(), logicNm, " => ", receiverArg.Data.(int))

	// if receiverArg.Data.(int) == 5 {
	// 	job.Action.Pause(receiverArg.State)
	// 	time.Sleep(10 * time.Second)
	// 	job.Action.Start(receiverArg.State)
	// }

}

func RunScheduler() {
	// f, err := os.Create("./dat2")
	// defer f.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered\n")
	// fmt.Printf("1\n", n4)
	// n5, err := w.WriteString("buffered\n")
	// fmt.Printf("2\n", n5)
	// w.Flush()

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
	// job.Action.Start(c)

	fmt.Println("XXX")
}
