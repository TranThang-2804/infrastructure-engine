package messagequeue

// MessageQueue defines the interface for a message queue system.
type MessageQueue interface {
	// Publish sends a message to the queue.
	Publish(message string) error

	// Subscribe registers a consumer to receive messages from the queue.
	// The consumer function is called for each message.
	Subscribe(consumer func(message string) error) error

	// Close shuts down the queue and cleans up resources.
	Close() error
}

