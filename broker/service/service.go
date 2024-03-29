package service

import (
  daprd "github.com/dapr/go-sdk/service/http"
  // daprd "github.com/dapr/go-sdk/service/grpc"
  "github.com/halokid/NeonRabbit/broker/handler"
  "github.com/halokid/NeonRabbit/unify"
)

type Service struct{}

func Run() error {

  // TODO: sub event
  // b.GetAdaptee().Sub()
  //go brokerx.BrokerxGlobal.Adaptee.Sub()

  s := daprd.NewService(unify.Unifyx.Pkg.Env.Broker.AppPort)
  // s, _ := daprd.NewService(unify.Unifyx.Pkg.Env.Broker.AppPort)

  // TODO: add handlers
  err := s.AddServiceInvocationHandler("/ping", handler.PingHandler)
  unify.Unifyx.Pkg.Logger.L.Infof("Broker ping handler err -->>> %+v", err)

  //err = s.AddServiceInvocationHandler("/pub", handler.PubHandler)
  //unify.Unifyx.Pkg.Logger.Infof("Broker pub handler err -->>> %+v", err)

  err = s.Start()
  unify.Unifyx.Pkg.Logger.L.Infof("Broker service err -->>> %+v", err)
  return nil
}
