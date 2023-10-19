package main

import (
  "context"
  "fmt"
  "log"
  "strconv"
  "time"

  "github.com/segmentio/kafka-go"
)

// the topic and broker address are initialized as constants
const (
  // topic = "message-log"
  topic = "neon-rabbit"
  //brokerAddress = "localhost:9092"
  brokerAddress = "192.168.1.129:9092"
  // broker1Address = "localhost:9093"
  // broker2Address = "localhost:9094"
  // broker3Address = "localhost:9095"
)

func produce(ctx context.Context) {
  // initialize a counter
  i := 0

  // intialize the writer with the broker addresses, and the topic
  w := kafka.NewWriter(kafka.WriterConfig{
    // Brokers: []string{broker1Address, broker2Address, broker3Address},
    Brokers: []string{brokerAddress},
    Topic:   topic,
  })

  for {
    // each kafka message has a key and value. The key is used
    // to decide which partition (and consequently, which broker)
    // the message gets published on
    // err := w.WriteMessages(ctx, kafka.Message{
    // 	Key: []byte(strconv.Itoa(i)),
    // 	// create an arbitrary message payload for the value
    // 	Value: []byte("this is message" + strconv.Itoa(i)),
    // })

    // TODO: just send message
    err := w.WriteMessages(ctx, kafka.Message{
      Value: []byte("this is message" + strconv.Itoa(i)),
    })

    if err != nil {
      panic("could not write message " + err.Error())
    }

    // log a confirmation once the message is written
    fmt.Println("writes:", i)
    i++
    // sleep for a second
    time.Sleep(time.Second)
  }
}

func consume(ctx context.Context) {
  // initialize a new reader with the brokers and topic
  // the groupID identifies the consumer and prevents
  // it from receiving duplicate messages
  r := kafka.NewReader(kafka.ReaderConfig{
    // Brokers: []string{broker1Address, broker2Address, broker3Address},
    Brokers: []string{brokerAddress},
    Topic:   topic,
    GroupID: "my-group",
  })

  log.Printf("total Offset -->>> %+v", r.Offset())

  for {
    // the `ReadMessage` method blocks until we receive the next event
    msg, err := r.ReadMessage(ctx)
    if err != nil {
      panic("could not read message " + err.Error())
    }
    // after receiving the message, log its value
    fmt.Println("received: ", string(msg.Value))
  }

}

func main() {
  // create a new context
  ctx := context.Background()

  // produce messages in a new go routine, since
  // both the produce and consume functions are
  // blocking

  // TODO: producer and consumer in the same context
  //go produce(ctx)
  consume(ctx)

  // TODO: only read the latest message from topic
  // consume(ctx)
}
