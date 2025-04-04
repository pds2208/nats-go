package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	sub, err := js.PullSubscribe("test", "i-am-durable", nats.PullMaxWaiting(128))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// Fetch will return as soon as any message is available rather than wait until the full batch size is
		// available, Using a batch size of more than 1 allows for higher throughput when needed.
		messages, err := sub.Fetch(10, nats.Context(ctx))
		if err != nil {
			log.Fatal(err)
		}

		for _, msg := range messages {
			fmt.Println("received msg ", string(msg.Data))
			err := msg.Ack()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
