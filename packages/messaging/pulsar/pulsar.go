package pulsar

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"

	"github.com/apache/pulsar-client-go/pulsar"
)

type pulsarClient struct {
	client    pulsar.Client
	producers []pulsar.Producer
	url       string
}

func NewPulsarClient(pulsarURL string) (messaging.Messaging, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsarURL,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	defer client.Close()
	if err != nil {
		return nil, err
	}
	return &pulsarClient{
		client: client,
		url:    pulsarURL,
	}, nil
}

func (pc pulsarClient) AddProducer(topicName string) error {
	producer, err := pc.client.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})

	pc.producers = append(pc.producers, producer)

	if err != nil {
		return err
	}
	return nil
}

func (pc pulsarClient) SendToTopic(topicName string, msg interface{}) error {
	for _, producer := range pc.producers {
		if producer.Topic() == topicName {
			marshalMsg, err := json.Marshal(msg)
			if err != nil {
				return err
			}
			_, err = producer.Send(
				context.Background(),
				&pulsar.ProducerMessage{
					Payload: marshalMsg,
				},
			)
			return nil
		}
	}
	return errors.New("not found producer for this topic name")
}
