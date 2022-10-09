package kafka

//IKafka ...
type IKafka interface {
	PushToKafkaStream(msg string)
}
