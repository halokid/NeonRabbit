package brokerx

import (
  "github.com/halokid/NeonRabbit/unify"
)

type Brokerx struct {
  Adaptee Adaptee
}

var BrokerxGlobal *Brokerx

func init() {
  BrokerxGlobal = NewBroker()
}

func NewBroker() *Brokerx {
  b := &Brokerx{}
  b.SetAdaptee()
  return b
}

func (b *Brokerx) SetAdaptee() error {
  // env, _ := pkg.ReadEnv()
  switch unify.Unifyx.Pkg.Env.Broker.Adapter {
  case "kafka":
    b.Adaptee = NewKafka()
  }
  return nil
}

func (b *Brokerx) GetAdaptee() Adaptee {
  return b.Adaptee
}
