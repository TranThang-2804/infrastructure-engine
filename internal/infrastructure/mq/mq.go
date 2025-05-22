package mq

// MessageQueue defines the interface for a message queue system.
type MessageQueue interface {
	// Publish sends a message to the queue.
	Publish(subject string, message []byte, opts ...any) error

	// Subscribe registers a consumer to receive messages from the queue.
	// The consumer function is called for each message.
	Subscribe(subject string, consumer func(message []byte) error) error

	// Close shuts down the queue and cleans up resources.
	Close() error
}
