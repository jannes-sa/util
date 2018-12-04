package main

import (
	"test/channel/cmd/ch1/logic1"
	"test/channel/cmd/ch1/logic2"
	"time"
)

func main() {
	logic1.RunScheduler()
	logic2.RunScheduler()

	time.Sleep(10 * time.Minute)
}
