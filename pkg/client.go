package pkg

import (
  "log"

  _daprClient "github.com/dapr/go-sdk/client"
)

var DaprClient _daprClient.Client

func initClient() {
  log.Println("-->>> Pkg DaprClient init")
  var err error
  DaprClient, err = _daprClient.NewClient()
  log.Printf("iniClient err -->>> %+v", err)
  if err != nil {
    panic(any(err))
  }
}
