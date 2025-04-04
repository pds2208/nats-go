package main

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = js.DeleteStream(ctx, "test-stream")

	_, _ = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "test-stream",
		Subjects: []string{"test"},
	})

	_ = nc.Drain()
	nc.Close()

}
