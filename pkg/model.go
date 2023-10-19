package pkg

type DBAdaptee int

const (
  Postgre  DBAdaptee = iota
  Mysql
  Sqlite
)

func (d DBAdaptee) String() string {
  return [...]string{"Postgre", "Mysql", "Sqlite"}[d]
}

var UseDbAdaptee DBAdaptee = Postgre