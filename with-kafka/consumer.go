package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	_const "kafka/const"
	"log"
	"time"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{_const.MyTopic, "^aRegex.*[Tt]opic"}, nil)

	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, msg.Value)
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v %v \n", err, msg)
		}
	}
}
