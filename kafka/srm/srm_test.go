package srm

import (
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	go Producer(5 * time.Second)
	Consumer()
	time.Sleep(5 * time.Minute)
}
