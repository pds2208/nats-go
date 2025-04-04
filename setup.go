package main

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"time"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = js.DeleteStream(ctx, "test-stream")

	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "test-stream",
		Subjects: []string{"test"},
	})

	if err != nil {
		log.Fatal(err)
	}

	_ = nc.Drain()
	nc.Close()

}
