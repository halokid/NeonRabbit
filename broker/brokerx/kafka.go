package brokerx

import (
  "context"

  "github.com/halokid/NeonRabbit/unify"
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
    Brokers:  []string{unify.Unifyx.Pkg.Env.Broker.Server},
    Topic:    unify.Unifyx.Pkg.Env.Broker.Topic,
    GroupID:  unify.Unifyx.Pkg.Env.Broker.GroupId,
    MinBytes: 10e1, // 0.1KB
    MaxBytes: 10e3, // 10KB
  }
}

func setWriterCfg() kafka.WriterConfig {
  return kafka.WriterConfig{
    Brokers: []string{unify.Unifyx.Pkg.Env.Broker.Server},
    Topic:   unify.Unifyx.Pkg.Env.Broker.Topic,
  }
}

func (k *Kafka) Sub() error {
  // initialize a new reader with the brokers and topic
  // the groupID identifies the consumer and prevents it
  // from receivig duplicate message
  unify.Unifyx.Pkg.Logger.Infof("Sub total Offset -->>> %+v", k.Reader.Offset())
  for {
    // the `ReadMessage` method blocks until we receive the next event
    msg, err := k.Reader.ReadMessage(context.Background())
    if err != nil {
      unify.Unifyx.Pkg.Logger.L.Errorf("-->>> Kafka Sub err: %+v", err)
    }
    unify.Unifyx.Pkg.Logger.L.Infof("-->>> Sub msg: %+v", string(msg.Value))
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
  unify.Unifyx.Pkg.Logger.Infof("-->>> Kafka pub err: %+v", err)

  return err
}
