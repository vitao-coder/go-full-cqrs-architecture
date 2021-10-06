package event

import (
	"context"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging"

	"github.com/vitao-coder/go-full-cqrs-architecture/components/cqrs"

	"github.com/pkg/errors"
)

type EventHandler interface {
	HandlerName() string
	NewEvent() interface{}
	Handle(ctx context.Context, event interface{}) error
}

type EventsSubscriberConstructor func(handlerName string) (messaging.Subscriber, error)

type EventProcessor struct {
	handlers              []EventHandler
	generateTopic         func(eventName string) string
	subscriberConstructor EventsSubscriberConstructor
	marshaller            cqrs.CommandEventMarshaller
	logger                logging.Logger
}

func NewEventProcessor(
	handlers []EventHandler,
	generateTopic func(eventName string) string,
	subscriberConstructor EventsSubscriberConstructor,
	marshaller cqrs.CommandEventMarshaller,
	logger logging.Logger,
) (*EventProcessor, error) {
	if len(handlers) == 0 {
		return nil, errors.New("missing handlers")
	}
	if generateTopic == nil {
		return nil, errors.New("nil generateTopic")
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

	return &EventProcessor{
		handlers,
		generateTopic,
		subscriberConstructor,
		marshaller,
		logger,
	}, nil
}

func (p EventProcessor) Handlers() []EventHandler {
	return p.handlers
}
