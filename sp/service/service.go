package service

import (
  daprd "github.com/dapr/go-sdk/service/http"
  // daprd "github.com/dapr/go-sdk/service/grpc"
  "github.com/halokid/NeonRabbit/sp/handler"
  "github.com/halokid/NeonRabbit/unify"
)

type Service struct{}

func Run() error {
  s := daprd.NewService(unify.Unifyx.Pkg.Env.Sp.AppPort)
  // s, _ := daprd.NewService(unify.Unifyx.Pkg.Env.Broker.AppPort)

  // TODO: add handlers
  err := s.AddServiceInvocationHandler("/ping", handler.PingHandler)
  unify.Unifyx.Pkg.Logger.L.Infof("Sp ping handler err -->>> %+v", err)

  err = s.AddServiceInvocationHandler("/user", handler.UserHandler)
  unify.Unifyx.Pkg.Logger.L.Infof("Sp User handler err -->>> %+v", err)

  err = s.Start()
  unify.Unifyx.Pkg.Logger.L.Infof("Sp service err -->>> %+v", err)
  return nil
}
