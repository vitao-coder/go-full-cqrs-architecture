package command

import (
	"context"
	"errors"

	"github.com/vitao-coder/go-full-cqrs-architecture/components/cqrs"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"
)

type CommandBus struct {
	publisher messaging.Publisher
	marshaler cqrs.CommandEventMarshaller
}

func NewCommandBus(
	publisher messaging.Publisher,
	marshaler cqrs.CommandEventMarshaller,
) (*CommandBus, error) {
	if publisher == nil {
		return nil, errors.New("missing publisher")
	}

	if marshaler == nil {
		return nil, errors.New("missing marshaller")
	}

	return &CommandBus{publisher, marshaler}, nil
}

func (c CommandBus) Send(ctx context.Context, cmd interface{}) error {
	msg, err := c.marshaler.Marshal(cmd)
	if err != nil {
		return err
	}

	commandName := c.marshaler.Name(cmd)

	err = c.publisher.CreateProducer(commandName)
	if err != nil {
		return err
	}
	topicName := commandName

	msg.SetContext(ctx)

	return c.publisher.Publish(topicName, msg)
}
