package adaptee

import (
	"context"
	"log"

	"github.com/halokid/NeonRabbit/pkg"
	kafka "github.com/segmentio/kafka-go"
)

type Kafka struct{}

func NewKafka() *Kafka {
	return &Kafka{}
}

func (k *Kafka) Sub(topic, groupId string) error {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents it
	// from receivig duplicate message
	r := kafka.NewReader(kafka.ReaderConfig{
		// Brokers: []string{broker1Address, broker2Address, broker3Address},
		Brokers: []string{pkg.EnvGlobal.Broker.Server},
		Topic:   topic,
		GroupID: groupId,
		MinBytes:  10e1, // 0.1KB
		MaxBytes:  10e3, // 10KB
		//MaxWait: 0,
	})

	log.Printf("total Offset -->>> %+v", r.Offset())
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(context.Background())
		pkg.Logger.Infof("-->>> Kafka sub err: %+v, msg: %+v", err, msg)
		pkg.Logger.Infof("-->>> Sub msg: %+v", string(msg.Value))
		// after receiving the message, do persist process
		// TODO:

	}

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
	pkg.Logger.Infof("-->>> Kafka pub err: %+v", err)

	return nil
}
