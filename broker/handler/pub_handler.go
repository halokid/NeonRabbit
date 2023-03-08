package handler

import (
	"context"
	"log"

	"github.com/dapr/go-sdk/service/common"
	"github.com/halokid/NeonRabbit/broker/brokerx"
	"github.com/halokid/NeonRabbit/unify"
)

func PubHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content,
	err error) {

	unify.Unifyx.Pkg.Logger.Info("-->>> call PubHandler")
	message := string(in.Data)
	unify.Unifyx.Pkg.Logger.Infof("PubHandler pub message -->>> %+v", message)
	brokerx.BrokerxGlobal.Adaptee.Pub(message)

	rsp := unify.Unifyx.Pkg.ConvertStructToByte(unify.Unifyx.Vo.SussceRsp())
	out = &common.Content{
		// Data: in.Data,
		Data:        rsp,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	log.Printf("PubHandler return out -->>> %+v, string data -->>> %+v",
		out, string(in.Data))

	return
}
