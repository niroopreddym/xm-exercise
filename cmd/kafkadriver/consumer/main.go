package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// func main() {

// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "localhost",
// 		"group.id":          "myGroup",
// 		"auto.offset.reset": "earliest",
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.SubscribeTopics([]string{"first_kafka_topic"}, nil)

// 	for {
// 		msg, err := c.ReadMessage(-1)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 		} else {
// 			// The client will automatically try to recover from all errors.
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 		}
// 	}
// }

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "first_kafka_topic", 0)
	if err != nil {
		fmt.Printf("Conn error: %v\n", err)
	}

	for {
		conn.SetDeadline(time.Now().Add(time.Second * 10))
		message, err := conn.ReadMessage(106)
		if err != nil {
			fmt.Printf("Conn error: %v\n", err)
		}
		fmt.Println(string(message.Value))
	}

}
