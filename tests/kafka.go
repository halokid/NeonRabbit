package main

import (
  "github.com/confluentinc/confluent-kafka-go/v2/kafka"
  "github.com/halokid/ColorfulRabbit"
  "github.com/halokid/NeonRabbit/pkg"
)

func Pub(topic, message string) error {
  p, err := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": pkg.EnvGlobal.Broker.Server})
  ColorfulRabbit.CheckError(err, "Kafka Pub error")
  deliveryChan := make(chan kafka.Event)
  err = p.Produce(&kafka.Message{
    TopicPartition: kafka.TopicPartition{
      Topic: &topic,
      Partition:  kafka.PartitionAny},
    Value:  []byte(message),
    Headers:  []kafka.Header{
      {
        Key:  "myTestHeader",
        Value:  []byte("header values are binary"),
      }}}, deliveryChan)

  e := <-deliveryChan
  m := e.(*kafka.Message)

  ColorfulRabbit.CheckError(m.TopicPartition.Error, "-->>> Kafka Pub error")
  return nil
}

func main() {
  err := Pub("neon_rabbit", "test pub")
  ColorfulRabbit.CheckError(err, "")
}








