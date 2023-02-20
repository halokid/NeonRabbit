package adaptee

import (
	"context"
	"time"

	"github.com/halokid/ColorfulRabbit"
	kafka "github.com/segmentio/kafka-go"
)

type Kafka struct{}

func NewKafka() *Kafka {
	return &Kafka{}
}

func (k *Kafka) Sub() error {
	return nil
}

func (k *Kafka) Pub(topic, message string) error {
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092",
		topic, partition)
	ColorfulRabbit.CheckError(err, "-->>> Kafka pub error")

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	ColorfulRabbit.CheckError(err, "-->>> Kafka pub err failed to write messages")

	err = conn.Close()
	ColorfulRabbit.CheckError(err, "-->>> Kafka pub failed to close writer")

	return nil
}
