package main

import (
  "context"
  "fmt"
  "github.com/segmentio/kafka-go"
  "log"
  "time"
)

func Consumer() error {
  topic := "neon-rabbit"
  partition := 0

  conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
  if err != nil {
    log.Fatal("failed to dial leader:", err)
  }

  conn.SetReadDeadline(time.Now().Add(10 * time.Second))
  // conn.ReadFirstOffset()
  batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

  b := make([]byte, 10e3) // 10KB max per message
  for {
    n, err := batch.Read(b)
    if err != nil {
      break
    }
    fmt.Println(string(b[:n]))
  }

  if err := batch.Close(); err != nil {
    log.Fatal("failed to close batch:", err)
  }

  if err := conn.Close(); err != nil {
    log.Fatal("failed to close connection:", err)
  }

  return nil
}

func main() {
  Consumer()
}
