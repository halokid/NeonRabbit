package brokerx

import (
	"context"

	"github.com/halokid/NeonRabbit/pkg"
	kafka "github.com/segmentio/kafka-go"
)

type Kafka struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

func NewKafka() *Kafka {
	reader := kafka.NewReader(setReaderCfg())
	writer := kafka.NewWriter(setWriterCfg())
	return &Kafka{
		Reader: reader,
		Writer: writer,
	}
}

func setReaderCfg() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers:  []string{pkg.Pkgx.Env.Broker.Server},
		Topic:    pkg.Pkgx.Env.Broker.Topic,
		GroupID:  pkg.Pkgx.Env.Broker.GroupId,
		MinBytes: 10e1, // 0.1KB
		MaxBytes: 10e3, // 10KB
	}
}

func setWriterCfg() kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers: []string{pkg.Pkgx.Env.Broker.Server},
		Topic:   pkg.Pkgx.Env.Broker.Topic,
	}
}

func (k *Kafka) Sub() error {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents it
	// from receivig duplicate message
	pkg.Pkgx.Logger.Infof("Sub total Offset -->>> %+v", k.Reader.Offset())
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := k.Reader.ReadMessage(context.Background())
		pkg.Pkgx.Logger.Infof("-->>> Kafka sub err: %+v, msg: %+v", err, msg)
		pkg.Pkgx.Logger.Infof("-->>> Sub msg: %+v", string(msg.Value))
		// after receiving the message, do persist process
		// TODO:
	}

	return nil
}

func (k *Kafka) Pub(message string) error {

	ctx := context.Background()
	err := k.Writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(message),
	})
	pkg.Pkgx.Logger.Infof("-->>> Kafka pub err: %+v", err)

	return err
}
