package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

const topic string = "example"
const messageInterval time.Duration = 500 * time.Millisecond

func main() {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	defer pubSub.Close()

	stopChan := make(chan bool)

	go subscribeTo(pubSub)
	go publishTo(pubSub, stopChan)

	waitForInterrupt()

	fmt.Println("Stopping application...")
	stopChan <- true
}

func subscribeTo(subscriber message.Subscriber) {
	messages, err := subscriber.Subscribe(context.Background(), topic)
	if err != nil {
		panic(err)
	}
	for message := range messages {
		fmt.Println("Received:", message.UUID, string(message.Payload))
		message.Ack()
	}
}

func publishTo(publisher message.Publisher, stopChan <-chan bool) {
	ticker := time.NewTicker(messageInterval)
	for {
		select {
		case t := <-ticker.C:
			msg := message.NewMessage(watermill.NewUUID(), []byte("Hi "+t.String()))
			if err := publisher.Publish(topic, msg); err != nil {
				fmt.Println("An error occured while publishing message", err)
			}
		case <-stopChan:
			return
		}
	}
}

func waitForInterrupt() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	<-interruptChan
}
