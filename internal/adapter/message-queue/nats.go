package mqadapter

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

// NatsSubject represents a NATS subject with an active subscription.
type NatsMQ struct {
	conn         *nats.Conn
	subjectNames []string
	sub          *nats.Subscription
	mu           sync.Mutex
}

// NewMessageQueue creates a new connection to the NATS server.
func NewNatsMQ(url string, subjectNames []string) (*NatsMQ, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsMQ{
		conn:         conn,
		subjectNames: subjectNames,
	}, nil
}

// Subscribe sets up a subscription to the subject with a message handler.
func (s *NatsMQ) Subscribe(subject string, handler func(message string) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sub != nil {
		return fmt.Errorf("already subscribed to subject: %s", subject)
	}

	sub, err := s.conn.Subscribe(subject, func(msg *nats.Msg) {
		if err := handler(string(msg.Data)); err != nil {
			fmt.Printf("Error processing message on '%s': %v\n", subject, err)
		}
	})
	if err != nil {
		return err
	}

	s.sub = sub
	return nil
}

// Publish sends a message to the subject.
func (s *NatsMQ) Publish(subject string, message string) error {
	return s.conn.Publish(subject, []byte(message))
}

// Close unsubscribes from the subject.
func (s *NatsMQ) Close() error {
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
