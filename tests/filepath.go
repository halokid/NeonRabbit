package main

import (
  "log"
  "runtime"
)

func RealPath() {
  a, b, c, d := runtime.Caller(1)
  log.Println("1: ", a)
  log.Println("2: ", b)
  log.Println("3: ", c)
  log.Println("4: ", d)
}
