package main

import (
	"github.com/halokid/NeonRabbit/broker"
)

func main() {
	b := broker.NewBroker()
	b.Run()
}
