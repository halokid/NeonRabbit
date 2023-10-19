package log

type Repository interface {
  InsertLog(logres string) error
  GetAll() *[]Log
}

