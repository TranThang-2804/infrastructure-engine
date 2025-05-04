package mqadapter

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

type NatsMQConnection struct {
	conn *nats.Conn
}

// NatsSubject represents a NATS subject with an active subscription.
type NatsSubject struct {
	name    string
	sub     *nats.Subscription
	mu      sync.Mutex
	mq      *NatsMQConnection
}

// NewMessageQueue creates a new connection to the NATS server.
func NewNatsMQConnection(url string) (*NatsMQConnection, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsMQConnection{conn: conn}, nil
}

// NewSubject creates a subject handler tied to this message queue.
func (mq *NatsMQConnection) NewSubject(name string) *NatsSubject {
	return &NatsSubject{
		name: name,
		mq:   mq,
	}
}

// Subscribe sets up a subscription to the subject with a message handler.
func (s *NatsSubject) Subscribe(handler func(message string) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sub != nil {
		return fmt.Errorf("already subscribed to subject: %s", s.name)
	}

	sub, err := s.mq.conn.Subscribe(s.name, func(msg *nats.Msg) {
		if err := handler(string(msg.Data)); err != nil {
			fmt.Printf("Error processing message on '%s': %v\n", s.name, err)
		}
	})
	if err != nil {
		return err
	}

	s.sub = sub
	return nil
}

// Publish sends a message to the subject.
func (s *NatsSubject) Publish(message string) error {
	return s.mq.conn.Publish(s.name, []byte(message))
}

// Close unsubscribes from the subject.
func (s *NatsSubject) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sub != nil {
		if err := s.sub.Unsubscribe(); err != nil {
			return err
		}
		s.sub = nil
	}
	return nil
}

