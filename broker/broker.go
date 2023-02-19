package broker

import (
  "github.com/halokid/NeonRabbit/broker/adaptee"
  "github.com/halokid/NeonRabbit/pkg"
)

type Broker struct { }

func NewBroker() Adaptee {
  env, _ := pkg.ReadEnv()
  switch env.Broker.Adapter {
  case "kafka":
    return adaptee.NewKafka()
  }
  return nil
}

