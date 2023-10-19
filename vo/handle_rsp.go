package vo

type HandleRsp struct {
  Code    int8        `json:"code"`
  Message string      `json:"message"`
  Data    interface{} `json:"data"`
}

func (v *Vo) SussceRsp() HandleRsp {
  return HandleRsp{
    Code:    0,
    Message: "success",
    Data:    nil,
  }
}

func (v *Vo) HandleRspVo(code int8, message string, data interface{}) HandleRsp {
  return HandleRsp{
    Code:    code,
    Message: message,
    Data:    data,
  }
}
