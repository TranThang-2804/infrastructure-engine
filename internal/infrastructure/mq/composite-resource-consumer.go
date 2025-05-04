package mq

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type compositeResourceConsumer struct {
	compositeResourceUsecase domain.CompositeResourceUsecase
	messageQueue             MessageQueue
}

func NewCompositeResourceConsumer(mq MessageQueue, compositeResourceUsecase domain.CompositeResourceUsecase) domain.CompositeResourceEventConsumer {
	return &compositeResourceConsumer{
		compositeResourceUsecase: compositeResourceUsecase,
		messageQueue:             mq,
	}
}

func (cc *compositeResourceConsumer) SubscribeToPendingSubject(c context.Context) error {
	// Publish message to queue
	cc.messageQueue.Subscribe("composite-resource.pending", cc.compositeResourceUsecase.HandlePending)
	return nil
}

func (cc *compositeResourceConsumer) SubscribeToProvisioningSubject(c context.Context) error {
	// Publish message to queue
	cc.messageQueue.Subscribe("composite-resource.provisioning", cc.compositeResourceUsecase.HandleProvisioning)
	return nil
}

func (cc *compositeResourceConsumer) SubscribeToDeletingSubject(c context.Context) error {
	// Publish message to queue
	cc.messageQueue.Subscribe("composite-resource.deleting", cc.compositeResourceUsecase.HandleDeleting)
	return nil
}
