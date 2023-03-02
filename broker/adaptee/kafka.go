package adaptee

import (
	"context"
	"fmt"
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
		Brokers:  []string{pkg.EnvGlobal.Broker.Server},
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e1, // 0.1KB
		MaxBytes: 10e3, // 10KB
		// RebalanceTimeout: time.Duration(1 * time.Second),
		// RebalanceTimeout: time.Duration(3 * time.Second),
		// RebalanceTimeout: time.Duration(100 * time.Microsecond),
		//MaxWait: 0,
		// GroupBalancers: []kafka.GroupBalancer{kafka.RoundRobinGroupBalancer{}},
		// GroupBalancers: []kafka.GroupBalancer{kafka.RangeGroupBalancer{}},
	})

	pkg.Logger.Infof("Sub total Offset -->>> %+v", r.Offset())
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
		Brokers:  []string{pkg.EnvGlobal.Broker.Server},
		Topic:    topic,
		//Balancer: &kafka.LeastBytes{},
		//Balancer: &kafka.RoundRobin{},
	})


	/*
	// TODO: if need more consumer, just set more topic
	for i := 0; i < 100; i++ {
		w.Topic = fmt.Sprintf("tp%d", i)
		ctx := context.Background()
		err := w.WriteMessages(ctx, kafka.Message{
			Value: []byte(message),
		})
		pkg.Logger.Infof("-->>> Kafka pub err: %+v", err)
	}
	 */

	ctx := context.Background()
	//err := w.WriteMessages(ctx, kafka.Message{
	//	Value: []byte(message),
	//})

	for i := 0; i < 100; i++ {
		err := w.WriteMessages(ctx, kafka.Message{
			//Value: []byte(message),
			Value: []byte(fmt.Sprintf("%s-%d", message, i)),
		})
		pkg.Logger.Infof("-->>> Kafka pub err: %+v", err)
	}

	//pkg.Logger.Infof("-->>> Kafka pub err: %+v", err)

	return nil
}



