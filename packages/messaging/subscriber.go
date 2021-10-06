package messaging

type Subscriber interface {
	Publish(topicName string, msg interface{}) error
	ClosePublisher(topicName string) error
}
