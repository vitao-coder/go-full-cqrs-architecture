package command

import (
	"context"
	"errors"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"
)

type CommandBus struct {
	publisher     messaging.Messaging
	generateTopic func(commandName string) string
	marshaler     CommandMarshaler
}

func NewCommandBus(
	publisher messaging.Messaging,
	generateTopic func(commandName string) string,
	marshaler CommandMarshaler,
) (*CommandBus, error) {
	if publisher == nil {
		return nil, errors.New("missing publisher")
	}
	if generateTopic == nil {
		return nil, errors.New("missing generateTopic")
	}
	if marshaler == nil {
		return nil, errors.New("missing marshaler")
	}

	return &CommandBus{publisher, generateTopic, marshaler}, nil
}

func (c CommandBus) SendCommand(ctx context.Context, cmd interface{}) error {
	msg, err := c.marshaler.Marshal(cmd)
	if err != nil {
		return err
	}

	commandName := c.marshaler.Name(cmd)
	topicName := c.generateTopic(commandName)

	msg.SetContext(ctx)

	return c.publisher.Publish(topicName, msg)
}
