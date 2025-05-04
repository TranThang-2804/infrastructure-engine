package queue

import (
	"errors"
	"github.com/nats-io/nats.go"
	"sync"
)

// NATSMessageQueue is an implementation of the MessageQueue interface using NATS.
type NATSMessageQueue struct {
	conn      *nats.Conn
	subject   string
	sub       *nats.Subscription
	mu        sync.Mutex
	messageIDs map[string]bool // Tracks acknowledged messages
}

// NewNATSMessageQueue creates a new NATSMessageQueue instance.
func NewNATSMessageQueue(url, subject string) (*NATSMessageQueue, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NATSMessageQueue{
		conn:      conn,
		subject:   subject,
		messageIDs: make(map[string]bool),
	}, nil
}

// Publish sends a message to the NATS queue.
func (mq *NATSMessageQueue) Publish(message string) error {
	return mq.conn.Publish(mq.subject, []byte(message))
}

// Subscribe registers a consumer to receive messages from the NATS queue.
func (mq *NATSMessageQueue) Subscribe(consumer func(message string) error) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if mq.sub != nil {
		return errors.New("subscription already exists")
	}

	sub, err := mq.conn.Subscribe(mq.subject, func(msg *nats.Msg) {
		// Process the message using the consumer function
		err := consumer(string(msg.Data))
		if err != nil {
			// Reject the message if the consumer fails
			mq.Reject(string(msg.Data), true)
		} else {
			// Acknowledge the message if the consumer succeeds
			mq.Acknowledge(string(msg.Data))
		}
	})
	if err != nil {
		return err
	}

	mq.sub = sub
	return nil
}

// Acknowledge confirms that a message has been successfully processed.
func (mq *NATSMessageQueue) Acknowledge(messageID string) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	// Mark the message as acknowledged
	mq.messageIDs[messageID] = true
	return nil
}

// Reject indicates that a message could not be processed and optionally requeues it.
func (mq *NATSMessageQueue) Reject(messageID string, requeue bool) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if !requeue {
		// Mark the message as acknowledged to prevent reprocessing
		mq.messageIDs[messageID] = true
	}
	// NATS does not natively support requeuing, so this is a no-op for now.
	return nil
}

// Close shuts down the queue and cleans up resources.
func (mq *NATSMessageQueue) Close() error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if mq.sub != nil {
		if err := mq.sub.Unsubscribe(); err != nil {
			return err
		}
		mq.sub = nil
	}

	mq.conn.Close()
	return nil
}
