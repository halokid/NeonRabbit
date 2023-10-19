package log

import "github.com/halokid/NeonRabbit/model"

func GetAll() *[]Log {
  var logs []Log
  model.DB.Select("id, logres, creattime").Find(&logs)
  return &logs
}

