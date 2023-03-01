package main

import (
	"github.com/halokid/NeonRabbit/broker"
	"github.com/halokid/NeonRabbit/pkg"
)

func main() {
	b := broker.NewBroker()
	// start Sub
	b.GetAdaptee().Sub(pkg.EnvGlobal.Broker.Topic, pkg.EnvGlobal.Broker.GroupId)
	b.Run()
}
