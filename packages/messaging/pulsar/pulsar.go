package pulsar

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

const defaultTimeout = 30 * time.Second

type pulsarClient struct {
	client    pulsar.Client
	producers []pulsar.Producer
	url       string
}

func NewPulsarClient(pulsarURL string) (*pulsarClient, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsarURL,
		OperationTimeout:  defaultTimeout,
		ConnectionTimeout: defaultTimeout,
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

func (pc pulsarClient) CreateProducer(topicName string) error {
	producer, err := pc.client.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})

	if err != nil {
		return err
	}
	pc.producers = append(pc.producers, producer)
	return nil
}

func (pc pulsarClient) Publish(topicName string, msg interface{}) error {
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

func (pc pulsarClient) ClosePublisher(topicName string) error {
	for _, producer := range pc.producers {
		if producer.Topic() == topicName {
			producer.Close()
		}
	}
	return errors.New("not found producer for this topic name")
}

func (pc pulsarClient) Subscribe(ctx context.Context, topicName string) (<-chan *interface{}, error) {
	return nil, nil
}

func (pc pulsarClient) CloseSubscriber(topicName string) error {
	return nil
}
