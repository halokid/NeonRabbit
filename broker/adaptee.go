package broker

type Adaptee interface {
	Sub(topic, groupId string) error
	Pub(topic, message string) error
}
