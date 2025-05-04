package messagequeue

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

type NATSMessageQueue struct {
	conn *nats.Conn
}

type NATSMessageQueueSubject struct {
	subject    string
	sub        *nats.Subscription
	mu         sync.Mutex
	messageIDs map[string]bool // Tracks acknowledged messages
}

// NewNATSMessageQueue creates a new NATS connection.
func NewNATSMessageQueue(url string) (*NATSMessageQueue, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NATSMessageQueue{conn: conn}, nil
}

// NewSubject creates a new subject for the given NATS connection.
func (mq *NATSMessageQueue) NewSubject(subject string) *NATSMessageQueueSubject {
	return &NATSMessageQueueSubject{
		subject:    subject,
		messageIDs: make(map[string]bool),
	}
}

// Subscribe registers a consumer for the subject.
func (s *NATSMessageQueueSubject) Subscribe(conn *nats.Conn, consumer func(message string) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	sub, err := conn.Subscribe(s.subject, func(msg *nats.Msg) {
		err := consumer(string(msg.Data))
		if err != nil {
			fmt.Printf("Failed to process message: %s\n", err)
		}
	})
	if err != nil {
		return err
	}

	s.sub = sub
	return nil
}

// Publish sends a message to the subject.
func (s *NATSMessageQueueSubject) Publish(conn *nats.Conn, message string) error {
	return conn.Publish(s.subject, []byte(message))
}

// Close unsubscribes from the subject.
func (s *NATSMessageQueueSubject) Close() error {
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
