package service

import (
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/halokid/NeonRabbit/broker/brokerx"
	"github.com/halokid/NeonRabbit/broker/handler"
	"github.com/halokid/NeonRabbit/pkg"
)

type Service struct{}

func Run() error {

	// TODO: sub event
	// b.GetAdaptee().Sub()
	go brokerx.BrokerxGlobal.Adaptee.Sub()

	s := daprd.NewService(pkg.Pkgx.Env.Broker.AppPort)

	// TODO: add handlers
	err := s.AddServiceInvocationHandler("/ping", handler.PingHandler)
	pkg.Pkgx.Logger.Infof("Broker ping handler err -->>> %+v", err)

	err = s.AddServiceInvocationHandler("/pub", handler.PubHandler)
	pkg.Pkgx.Logger.Infof("Broker pub handler err -->>> %+v", err)

	err = s.Start()
	pkg.Pkgx.Logger.Infof("Broker service err -->>> %+v", err)
	return nil
}
