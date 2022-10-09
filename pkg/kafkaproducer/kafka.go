package kafka

import (
	"context"
	"fmt"
	"time"

	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/segmentio/kafka-go"
)

//XMKafka provides implementation of kafkaIface
type XMKafka struct {
}

//NewXMKafka is ctor
func NewXMKafka() IKafka {
	return &XMKafka{}
}

//PushToKafkaStream ...
func (xmakfka *XMKafka) PushToKafkaStream(msg string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "first_kafka_topic", 0)
	if err != nil {
		fmt.Printf("Conn error: %v\n", err)
	}

	conn.SetDeadline(time.Now().Add(time.Second * 10))
	n, err := conn.WriteMessages(kafka.Message{
		Value: []byte(
			msg,
		),
	})
	if err != nil {
		fmt.Printf("Conn error: %v\n", err)
	}

	fmt.Println("msgs written", n)
}
