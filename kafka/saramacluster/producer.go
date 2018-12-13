package saramacluster

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func Producer(intervalPublishMsg time.Duration) {

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		i := 1
		for t := range time.Tick(intervalPublishMsg) {
			msg := "message " + strconv.Itoa(i)
			log.Println("producer sent msg", t)
			sendInput(producer.Input(), msg)
			i++
		}
	}()

	monitoring(producer)

}

func sendInput(input chan<- *sarama.ProducerMessage, msg string) {
	input <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder(msg)}
}

func monitoring(producer sarama.AsyncProducer) {
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// var enqueued, errors int
	// var errors int
	// ProducerLoop:
	for {
		select {
		// case producer.Input() <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder("testing 123")}:
		// 	enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
		// errors++
		case <-signals:
			log.Println("STOP")
			return
		}
	}

	// log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
