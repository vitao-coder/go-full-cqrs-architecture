package messaging

type Messaging interface {
	CreateProducer(topicName string) error
	Publish(topicName string, msg interface{}) error
	Close(topicName string) error
}
