package model

import (
  "github.com/halokid/NeonRabbit/pkg"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)

var DB *gorm.DB

func init() {
  switch pkg.UseDbAdaptee {
  case pkg.Postgre:
    log.Println("-->>> Use Postgre DB")
    postgreConn()
  case pkg.Mysql:
    log.Println("-->>> Use Mysql DB")
  }
}

func postgreConn() (*gorm.DB, error) {
  DB, err := gorm.Open(postgres.New(postgres.Config{
    DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
    PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})
  log.Println("postgres err -->>>", err)
  return DB, err
}

