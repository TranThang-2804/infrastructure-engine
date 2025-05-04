package mq

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
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
	cr.messageQueue.Publish("composite-resource.pending", compositeResource.Id)
	return nil
}

func (cr *CompositeResourcePublisher) PublishToProvisioningSubject(c context.Context, compositeResource domain.CompositeResource) error {
	// Publish message to queue
	cr.messageQueue.Publish("composite-resource.provisioning", compositeResource.Id)
	return nil
}

func (cr *CompositeResourcePublisher) PublishToDeletingSubject(c context.Context, compositeResource domain.CompositeResource) error {
	// Publish message to queue
	cr.messageQueue.Publish("composite-resource.deleting", compositeResource.Id)
	return nil
}
