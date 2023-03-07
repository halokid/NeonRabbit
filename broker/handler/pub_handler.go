package handler

import (
	"context"
	"log"

	"github.com/dapr/go-sdk/service/common"
	"github.com/halokid/NeonRabbit/broker/brokerx"
	"github.com/halokid/NeonRabbit/pkg"
)

func PubHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content,
	err error) {

	pkg.Pkgx.Logger.Info("-->>> call PubHandler")
	message := string(in.Data)
	pkg.Pkgx.Logger.Infof("PubHandler pub message -->>> %+v", message)
	brokerx.BrokerxGlobal.Adaptee.Pub(message)

	rsp := pkg.ConvertStructToByte(pkg.Pkgx.Vo.SussceRsp())
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
