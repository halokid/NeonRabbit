package pkg

import (
  "bytes"
  "encoding/json"
)

func (p *Pkg) ConvertStructToByte(s interface{}) []byte {
  rsp := new(bytes.Buffer)
  err := json.NewEncoder(rsp).Encode(s)
  if err != nil {
    panic(err)
  }
  return rsp.Bytes()
}
