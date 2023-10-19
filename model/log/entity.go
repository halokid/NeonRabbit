package log

type Log struct {
  Id      int    `json:"id" gorm:"column:id;primary_key"`
  Logres  string `json:"logres"`
  Createtime  string  `json:"createtime"`
}
