package mq

import (
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

func (cc *compositeResourceConsumer) StartConsumer() error {
	// Subscribe to the pending subject
	if err := cc.subscribeToPendingSubject(); err != nil {
		return err
	}
	// Subscribe to the provisioning subject
	if err := cc.subscribeToProvisioningSubject(); err != nil {
		return err
	}
	// Subscribe to the deleting subject
	if err := cc.subscribeToDeletingSubject(); err != nil {
		return err
	}
	return nil
}

func (cc *compositeResourceConsumer) subscribeToPendingSubject() error {
	// Publish message to queue
	return cc.messageQueue.Subscribe("composite-resource.pending", cc.compositeResourceUsecase.HandlePending)
}

func (cc *compositeResourceConsumer) subscribeToProvisioningSubject() error {
	// Publish message to queue
	return cc.messageQueue.Subscribe("composite-resource.provisioning", cc.compositeResourceUsecase.HandleProvisioning)
}

func (cc *compositeResourceConsumer) subscribeToDeletingSubject() error {
	// Publish message to queue
	return cc.messageQueue.Subscribe("composite-resource.deleting", cc.compositeResourceUsecase.HandleDeleting)
}
