package brokerx

type Adaptee interface {
  Sub() error
  Pub(message string) error
}
