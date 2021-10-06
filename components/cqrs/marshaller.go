package cqrs

import (
	"github.com/vitao-coder/go-full-cqrs-architecture/components/cqrs/message"
)

type CommandEventMarshaller interface {
	Marshal(v interface{}) (*message.Message, error)
	Unmarshal(msg *message.Message, v interface{}) (err error)
	Name(v interface{}) string
	NameFromMessage(msg *message.Message) string
}
