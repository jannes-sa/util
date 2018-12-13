package saramacluster

import (
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	go Producer(1 * time.Second)
	Consumer()
	time.Sleep(5 * time.Minute)
}

func TestWhenConsumerCollectAfterProducerPublishSeveral(t *testing.T) {
	go Producer(1 * time.Second)
	time.Sleep(5 * time.Second)
	Consumer()
	time.Sleep(5 * time.Minute)
}
