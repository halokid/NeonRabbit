package handler

import (
  "context"
  "errors"
  "log"

  "github.com/dapr/go-sdk/service/common"
  "github.com/halokid/NeonRabbit/unify"
)

func UserHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content,
  err error) {
  unify.Unifyx.Pkg.Logger.L.Infof("-->>> call UserHandler, `in` is -->>> %+v", string(in.Data))
  if in == nil {
    err = errors.New("invocation parameter required")
    return
  }

  unify.Unifyx.Pkg.Logger.L.Infof(
    "user - ContentType:%s, Verb:%s, QueryString:%s, Body: %s",
    in.ContentType, in.Verb, in.QueryString, in.Data,
  )
  out = &common.Content{
    Data:        in.Data,
    ContentType: in.ContentType,
    DataTypeURL: in.DataTypeURL,
  }

  unify.Unifyx.Pkg.Logger.L.Infof("UserHandler return out -->>> %+v, string data -->>> %+v",
    out, string(in.Data))
  log.Printf("%+v", string(in.Data))
  return
}
