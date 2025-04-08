package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"strconv"
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

	_, err = js.Stream(ctx, "test-stream")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		_, err := js.PublishAsync("test", []byte("hello message "+strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Published hello message %d\n", i)
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
}
