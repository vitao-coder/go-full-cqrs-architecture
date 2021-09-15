package messaging

type Messaging interface {
	AddProducer(topicName string) error
	SendToTopic(topicName string, msg interface{}) error
}
