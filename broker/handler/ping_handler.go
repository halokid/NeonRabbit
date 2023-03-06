package handler

import (
	"context"
	"errors"
	"log"

	"github.com/dapr/go-sdk/service/common"
)

func PingHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content,
	err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"ping - ContentType:%s, Verb:%s, QueryString:%s, Body: %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	log.Printf("PingHandler return out -->>> %+v, string data -->>> %+v",
		out, string(in.Data))
	return
}
