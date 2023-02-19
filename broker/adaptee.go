package broker

type Adaptee interface {
  Sub() error
  Pub(topic, message string) error
}




