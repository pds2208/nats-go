package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	// Timeout after 280 odd years
	// can use a context instead but just to showcase
	const maxDuration time.Duration = 1<<63 - 1

	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	sub, err := js.PullSubscribe(
		"test",
		"i-am-durable",
		nats.PullMaxWaiting(128))

	if err != nil {
		log.Fatal(err)
	}

	for {

		// Fetch will return as soon as any message is available rather than wait until the full batch size is
		// available, Using a batch size of more than 1 allows for higher throughput when needed.
		// Waiting for 280 odd years for messages to arrive!
		messages, err := sub.Fetch(10, nats.MaxWait(maxDuration))
		if err != nil {
			log.Fatal(err)
		}

		for _, msg := range messages {
			fmt.Println("received msg ", string(msg.Data))
			err := msg.Ack()
			if err != nil {
				log.Fatal(err) // just crash if there is an error for showcase
			}
		}
	}
}
