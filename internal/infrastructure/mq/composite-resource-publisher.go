package mq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
)

type CompositeResourcePublisher struct {
	messageQueue MessageQueue
}

func NewCompositeResourcePublisher(mq MessageQueue) domain.CompositeResourceEventPublisher {
	return &CompositeResourcePublisher{
		messageQueue: mq,
	}
}

func (cr *CompositeResourcePublisher) PublishToPendingSubject(c context.Context, compositeResource domain.CompositeResource) error {
	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}
	cr.messageQueue.Publish("composite-resource.pending", messageData)
	log.BaseLogger.Info("Publish to pending subject", "compositeResourceId", compositeResource.Id)
	return nil
}

func (cr *CompositeResourcePublisher) PublishToProvisioningSubject(c context.Context, compositeResource domain.CompositeResource) error {
	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}
	return cr.messageQueue.Publish("composite-resource.provisioning", messageData)
}

func (cr *CompositeResourcePublisher) PublishToProvisioningSubjectWithDelay(c context.Context, compositeResource domain.CompositeResource) error {
	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}

	delay := 15 * time.Second
	return cr.messageQueue.PublishAfterDelay("composite-resource.provisioning", messageData, delay)
}

func (cr *CompositeResourcePublisher) PublishToDeletingSubject(c context.Context, compositeResource domain.CompositeResource) error {
	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		log.BaseLogger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}
	return cr.messageQueue.Publish("composite-resource.deleting", messageData)
}
