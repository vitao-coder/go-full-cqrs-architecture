package messaging

type Publisher interface {
	CreateProducer(topicName string) error
	Publish(topicName string, msg interface{}) error
	ClosePublisher(topicName string) error
}
