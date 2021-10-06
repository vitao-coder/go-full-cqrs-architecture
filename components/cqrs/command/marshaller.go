package command

import "github.com/vitao-coder/go-full-cqrs-architecture/components/message"

type CommandMarshaler interface {
	Marshal(v interface{}) (*message.Message, error)
	Unmarshal(msg *message.Message, v interface{}) (err error)
	Name(v interface{}) string
	NameFromMessage(msg *message.Message) string
}
