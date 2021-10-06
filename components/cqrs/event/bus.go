package event

import (
	"context"
	"errors"

	"github.com/vitao-coder/go-full-cqrs-architecture/components/cqrs"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"
)

type EventBus struct {
	publisher messaging.Publisher
	marshaler cqrs.CommandEventMarshaller
}

func NewEventBus(
	publisher messaging.Publisher,
	marshaler cqrs.CommandEventMarshaller,
) (*EventBus, error) {
	if publisher == nil {
		return nil, errors.New("missing publisher")
	}
	if marshaler == nil {
		return nil, errors.New("missing marshaller")
	}

	return &EventBus{publisher, marshaler}, nil
}

func (c EventBus) Publish(ctx context.Context, event interface{}) error {
	msg, err := c.marshaler.Marshal(event)
	if err != nil {
		return err
	}

	eventName := c.marshaler.Name(event)
	err = c.publisher.CreateProducer(eventName)
	if err != nil {
		return err
	}
	topicName := eventName

	msg.SetContext(ctx)

	return c.publisher.Publish(topicName, msg)
}
