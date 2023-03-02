package main

import (
	"github.com/halokid/NeonRabbit/broker"
	"github.com/halokid/NeonRabbit/pkg"
)

func main() {
	b := broker.NewBroker()
	// start Sub
	b.GetAdaptee().Sub(pkg.EnvGlobal.Broker.Topic, "neon-rabbit-group")
	//pkg.EnvGlobal.Broker.GroupId)
	// b.GetAdaptee().Sub(pkg.EnvGlobal.Broker.Topic, "neon-rabbit-group2")

	// b.Run()
}
