package mq

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/nats-io/nats.go"
)

// NatsMQ represents a NATS JetStream-based message queue.
type NatsMQ struct {
	conn         *nats.Conn
	js           nats.JetStreamContext
	subjectNames []string
	mu           sync.Mutex
	subs         map[string]*nats.Subscription
}

// NewNatsMQ initializes the JetStream context and connects to NATS.
func NewNatsMQ(url string, subjectNames []string) (MessageQueue, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	js, err := conn.JetStream()
	if err != nil {
		conn.Close()
		return nil, err
	}
	//
	// Define the stream configuration
	streamConfig := &nats.StreamConfig{
		Name:     "COMPOSITE_RESOURCE_EVENTS",      // Stream name
		Subjects: []string{"composite-resource.*"}, // Subject this stream listens to
		Storage:  nats.FileStorage,                 // Storage type
	}

	// Check if the stream exists
	_, err = js.StreamInfo(streamConfig.Name)
	if err == nil {
		// Stream exists, no need to create it again
		log.Logger.Info("Stream already exists. Skipping creation.")
	} else if err != nats.ErrStreamNotFound {
		// Some other error occurred
		log.Logger.Fatal("Jetstream queue has some error", "error", err)
	} else {
		// Stream does not exist, create it
		_, err = js.AddStream(streamConfig)
		if err != nil {
			log.Logger.Fatal("Cannot create Jetstream queue", "error", err)
		}
		log.Logger.Info("Jetstream queue created successfully")
	}

	return &NatsMQ{
		conn:         conn,
		js:           js,
		subjectNames: subjectNames,
		subs:         make(map[string]*nats.Subscription),
	}, nil
}

// Subscribe uses JetStream to create a durable consumer with manual ack and ack wait.
func (mq *NatsMQ) Subscribe(subject string, handler func(message []byte) error) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if _, exists := mq.subs[subject]; exists {
		return fmt.Errorf("already subscribed to subject: %s", subject)
	}

	sub, err := mq.js.Subscribe(subject, func(msg *nats.Msg) {
		if err := handler(msg.Data); err != nil {
			log.Logger.Error("❌ Error processing message on '%s': %v\n", subject, err)
			// Don't ack to trigger retry after AckWait
			return
		}
		log.Logger.Debug("Message handled successful")
		msg.Ack()
	}, nats.Durable("worker-"+strings.ReplaceAll(subject, ".", "-")),
		nats.ManualAck(),
		nats.AckWait(30*time.Second), // Visibility timeout
		nats.MaxDeliver(5),           // Max retry attempts
	)
	if err != nil {
		return err
	}

	mq.subs[subject] = sub
	return nil
}

// Publish sends a message using JetStream (persistent).
func (mq *NatsMQ) Publish(subject string, message []byte, opts ...interface{}) error {
	// Convert opts to []nats.PubOpt
	var pubOpts []nats.PubOpt
	for _, opt := range opts {
		if pubOpt, ok := opt.(nats.PubOpt); ok {
			pubOpts = append(pubOpts, pubOpt)
		} else {
			return fmt.Errorf("invalid publish option type: %T", opt)
		}
	}

	// Call the JetStream Publish method
	_, err := mq.js.Publish(subject, message, pubOpts...)
	return err
}

// Close unsubscribes from all subjects and closes the connection.
func (mq *NatsMQ) Close() error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	for subject, sub := range mq.subs {
		if err := sub.Unsubscribe(); err != nil {
			return fmt.Errorf("failed to unsubscribe from %s: %w", subject, err)
		}
	}
	mq.subs = nil
	mq.conn.Drain() // Close gracefully
	return nil
}
