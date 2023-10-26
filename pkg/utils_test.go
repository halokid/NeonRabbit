package pkg

import "testing"

func Test_ConConvertStructToByte(t *testing.T) {
  type rsp struct {
    Code int
  }

  r := rsp{
    Code: 9,
  }
  p := Pkg{}
  rspB := p.ConvertStructToByte(r)
  t.Logf("rspB -->>> %+v", rspB)
}
