package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        "localhost:9092",
		"group.id":                 "myGroup",
		"enable.auto.offset.store": false,
		"auto.offset.reset":        "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"foo"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			time.Sleep(3000 * time.Millisecond)
			c.CommitMessage(msg)
			fmt.Println("--->  commit message")
			time.Sleep(3000 * time.Millisecond)
			//fmt.Println(abc)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
