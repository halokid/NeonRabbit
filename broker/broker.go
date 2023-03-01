package broker

import (
	"github.com/halokid/NeonRabbit/broker/adaptee"
	"github.com/halokid/NeonRabbit/pkg"
)

type Broker struct{}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) GetAdaptee() Adaptee {
	env, _ := pkg.ReadEnv()
	switch env.Broker.Adapter {
	case "kafka":
		return adaptee.NewKafka()
	}
	return nil
}

func (b *Broker) Run() error {

	brokerAdapter := b.GetAdaptee()
	prefix := "neon-rabbit"
	// run Sub as daemon
	go brokerAdapter.Sub(prefix, prefix + "-group")

	// run Pub for service call


	return nil
}
