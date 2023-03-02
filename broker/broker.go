package broker

import (
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/halokid/NeonRabbit/broker/adaptee"
	"github.com/halokid/NeonRabbit/broker/handler"
	"github.com/halokid/NeonRabbit/pkg"
)

type Broker struct{}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) GetAdaptee() Adaptee {
	// env, _ := pkg.ReadEnv()
	switch pkg.EnvGlobal.Broker.Adapter {
	case "kafka":
		return adaptee.NewKafka()
	}
	return nil
}

func (b *Broker) Run() error {

	b.GetAdaptee().Sub(pkg.EnvGlobal.Broker.Topic, pkg.EnvGlobal.Broker.GroupId)

	// run Broker service
	s := daprd.NewService(":" + pkg.EnvGlobal.Broker.AppPort)

	// add handlers
	err := s.AddServiceInvocationHandler("/ping", handler.PingHandler)
	pkg.Logger.Infof("Broker ping handler err -->>> %+v", err)

	err = s.Start()
	pkg.Logger.Infof("Broker service err -->>> %+v", err)
	return err
}
