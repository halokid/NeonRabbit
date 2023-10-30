package handler

import (
  "context"
  "errors"
  "fmt"
  "log"

  dapr "github.com/dapr/go-sdk/client"
  "github.com/dapr/go-sdk/service/common"
  "github.com/halokid/NeonRabbit/pkg"
  "github.com/halokid/NeonRabbit/unify"
)

func JobHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content,
  err error) {

  unify.Unifyx.Pkg.Logger.L.Infof("-->>> call JobHandler, `in` is -->>> %+v", string(in.Data))
  if in == nil {
    err = errors.New("invocation parameter required")
    return
  }

  // TODO: call sp service user method
  //ctx := context.Background()
  /*
    client, err := dapr.NewClient()
    log.Printf("Job handler dapr celint -->>> %+v", client)
    if err != nil {
      panic(any(err))
    }
    defer client.Close()
  */
  content := &dapr.DataContent{
    ContentType: "text/plain",
    Data:        []byte(`{"userId":777}`),
  }
  resp, err := pkg.DaprClient.InvokeMethodWithContent(ctx, "neon_sp", "user", "post", content)
  if err != nil {
    panic(any(err))
  }
  fmt.Printf("service method invoked, response -->>> %s", string(resp))

  unify.Unifyx.Pkg.Logger.L.Infof(
    "job - ContentType:%s, Verb:%s, QueryString:%s, Body: %s",
    in.ContentType, in.Verb, in.QueryString, in.Data,
  )
  out = &common.Content{
    Data:        in.Data,
    ContentType: in.ContentType,
    DataTypeURL: in.DataTypeURL,
  }

  unify.Unifyx.Pkg.Logger.L.Infof("JobHandler return out -->>> %+v, string data -->>> %+v",
    out, string(out.Data))
  log.Printf("%+v", string(in.Data))
  return
}
