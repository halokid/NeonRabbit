package service

import (
  //daprd "github.com/dapr/go-sdk/service/http"
  "log"
  "net"

  daprd "github.com/dapr/go-sdk/service/grpc"
  "github.com/halokid/NeonRabbit/broker/handler"
  "github.com/halokid/NeonRabbit/unify"
)

type Service struct{}

func Run() error {

  /*
  // TODO: HTTP
  // b.GetAdaptee().Sub()
  //go brokerx.BrokerxGlobal.Adaptee.Sub()

  s := daprd.NewService(unify.Unifyx.Pkg.Env.Broker.AppPort)
  // s, _ := daprd.NewService(unify.Unifyx.Pkg.Env.Broker.AppPort)

  err := s.AddServiceInvocationHandler("/ping", handler.PingHandler)
  unify.Unifyx.Pkg.Logger.L.Infof("Broker ping handler err -->>> %+v", err)

  //err = s.AddServiceInvocationHandler("/pub", handler.PubHandler)
  //unify.Unifyx.Pkg.Logger.Infof("Broker pub handler err -->>> %+v", err)

  err = s.Start()
  unify.Unifyx.Pkg.Logger.L.Infof("Broker service err -->>> %+v", err)
  */

  ///*
  // TODO: GRPC 
  // Create a new gRPC service
  s, err := daprd.NewService(unify.Unifyx.Pkg.Env.Broker.AppPort)
  unify.Unifyx.Pkg.Logger.L.Infof("Broker ping handler err -->>> %+v", err)
  if err != nil {
    log.Fatalf("failed to start the server: %v", err)
    return err
  }

  // Add a service invocation handler
  err = s.AddServiceInvocationHandler("ping", handler.PingHandler)
  if err != nil {
    log.Fatalf("error adding service invocation handler: %v", err)
    return  err
  }

  // Start the gRPC service
  log.Println("Starting gRPC service...")
  if err := s.Start(); err != nil && err != net.ErrClosed {
    log.Fatalf("error starting service: %v", err)
  }
  //*/

  return  nil
}




