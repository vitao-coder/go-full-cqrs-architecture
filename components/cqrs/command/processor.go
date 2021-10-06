package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging"

	"github.com/vitao-coder/go-full-cqrs-architecture/components/cqrs"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"
)

type DuplicateCommandHandlerError struct {
	CommandName string
}

func (d DuplicateCommandHandlerError) Error() string {
	return fmt.Sprintf("command handler for command %s already exists", d.CommandName)
}

type CommandHandler interface {
	HandlerName() string
	NewCommand() interface{}
	Handle(ctx context.Context, cmd interface{}) error
}

type CommandsSubscriberConstructor func(handlerName string) (messaging.Publisher, error)

type CommandProcessor struct {
	handlers              []CommandHandler
	generateTopic         func(commandName string) string
	subscriberConstructor CommandsSubscriberConstructor
	marshaller            cqrs.CommandEventMarshaller
	logger                logging.Logger
}

func NewCommandProcessor(
	handlers []CommandHandler,
	generateTopic func(commandName string) string,
	subscriberConstructor CommandsSubscriberConstructor,
	marshaller cqrs.CommandEventMarshaller,
	logger logging.Logger,
) (*CommandProcessor, error) {
	if len(handlers) == 0 {
		return nil, errors.New("missing handlers")
	}
	if generateTopic == nil {
		return nil, errors.New("missing generateTopic")
	}
	if subscriberConstructor == nil {
		return nil, errors.New("missing subscriberConstructor")
	}
	if marshaller == nil {
		return nil, errors.New("missing marshaller")
	}
	if logger == nil {
		return nil, errors.New("missing logger")
	}

	return &CommandProcessor{
		handlers,
		generateTopic,
		subscriberConstructor,
		marshaller,
		logger,
	}, nil
}

func (p CommandProcessor) Handlers() []CommandHandler {
	return p.handlers
}
