package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"strconv"
	"time"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	js.Stream(ctx, "test-stream")

	for i := 0; i < 100; i++ {
		js.PublishAsync("test", []byte("hello message "+strconv.Itoa(i)))
		fmt.Printf("Published hello message %d\n", i)
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
}
