package handler

import (
    "context"
    "github.com/dapr/go-sdk/service/common"
    "github.com/halokid/NeonRabbit/broker/brokerx"
    "github.com/halokid/NeonRabbit/pkg"
)

func PubHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content,
    err error) {

    pkg.Logger.Info("-->>> call PubHandler")
    message := string(in.Data)
    pkg.Logger.Infof("PubHandler pub message -->>> %+v", message)
    brokerx.BrokerxGlobal.Adaptee.Pub(message)
    return
}
