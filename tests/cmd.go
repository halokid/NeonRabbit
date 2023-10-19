package main

import (
  dapr "github.com/dapr/go-sdk/client"
)

func main() {
  // --- kafka ---
  // log.Println(time.Now().Unix())
  // go Consumer()
  // time.Sleep(5 * time.Second)

  // log.Println("-->>> Satrt consumer done.")
  // Producer() // TODO: reach here program will quit
  // time.Sleep(100 * time.Second)
  // log.Println(time.Now().Unix())

  // --- zap ---
  //ZapC1()
  // ZapC2()
  // Zapc3()

  // RealPath()

  // -------------------------- DAPR -------------------------
  client, err := dapr.NewClient()
  if err != nil {
    panic(err)
  }
  defer client.Close()
  // TODO: use the client here, see below for examples
}
