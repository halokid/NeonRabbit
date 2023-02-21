package adaptee

import (
	"context"

	"github.com/halokid/NeonRabbit/pkg"
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

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{pkg.EnvGlobal.Broker.Server},
		Topic:   topic,
	})
	ctx := context.Background()
	err := w.WriteMessages(ctx, kafka.Message{
		Value: []byte(message),
	})
	pkg.Logger.Errorf("-->>> Kafka pub err: %+v", err)

	return nil
}
