package message

import (
	"bytes"
	"context"
	"sync"
)

type Payload []byte

type ackType int

const (
	noAckSent ackType = iota
	ack
	nack
)

var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type Message struct {
	GUID string

	Metadata Metadata

	Payload Payload

	ack   chan struct{}
	noAck chan struct{}

	ackMutex    sync.Mutex
	ackSentType ackType
	ctx         context.Context
}

func NewMessage(guid string, payload Payload) *Message {
	return &Message{
		GUID:     guid,
		Metadata: make(map[string]string),
		Payload:  payload,
		ack:      make(chan struct{}),
		noAck:    make(chan struct{}),
	}
}

func (m *Message) Equals(toCompare *Message) bool {
	if m.GUID != toCompare.GUID {
		return false
	}
	if len(m.Metadata) != len(toCompare.Metadata) {
		return false
	}
	for key, value := range m.Metadata {
		if value != toCompare.Metadata[key] {
			return false
		}
	}
	return bytes.Equal(m.Payload, toCompare.Payload)
}

func (m *Message) Ack() bool {
	m.ackMutex.Lock()
	defer m.ackMutex.Unlock()

	if m.ackSentType == nack {
		return false
	}
	if m.ackSentType != noAckSent {
		return true
	}

	m.ackSentType = ack
	if m.ack == nil {
		m.ack = closedchan
	} else {
		close(m.ack)
	}

	return true
}

func (m *Message) Nack() bool {
	m.ackMutex.Lock()
	defer m.ackMutex.Unlock()

	if m.ackSentType == ack {
		return false
	}
	if m.ackSentType != noAckSent {
		return true
	}

	m.ackSentType = nack

	if m.noAck == nil {
		m.noAck = closedchan
	} else {
		close(m.noAck)
	}

	return true
}

func (m *Message) Acked() <-chan struct{} {
	return m.ack
}

func (m *Message) Nacked() <-chan struct{} {
	return m.noAck
}

func (m *Message) Context() context.Context {
	if m.ctx != nil {
		return m.ctx
	}
	return context.Background()
}

func (m *Message) SetContext(ctx context.Context) {
	m.ctx = ctx
}

func (m *Message) Copy() *Message {
	msg := NewMessage(m.GUID, m.Payload)
	for k, v := range m.Metadata {
		msg.Metadata.Set(k, v)
	}
	return msg
}
