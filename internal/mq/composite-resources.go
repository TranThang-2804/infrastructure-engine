package mq

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/message-queue"
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
)

func NewCompositeResourceMQController(gitStore git.GitStore, timeout time.Duration) {
	cr := repository.NewCompositeResourceRepository(gitStore)
	br := repository.NewBluePrintRepository(gitStore)
	bu := usecase.NewBluePrintUsecase(br, timeout)
	cp := &controller.CompositeResourceController{
		CompositeResourceUseCase: usecase.NewCompositeResourceUsecase(cr, bu, timeout),
	}

	// Create a NATS connection
	mq, err := mqadapter.NewNatsMQConnection(env.Env.NATS_URL)
	if err != nil {
		log.Logger.Fatal("Failed to connect to NATS", "error", err)
	}

	// Create a subject
	compositeResourcePendingSubject := mq.NewSubject("composite-resource.pending")
	compositeResourceProvisioningSubject := mq.NewSubject("composite-resource.provisioning")
	compositeResourceDeletingSubject := mq.NewSubject("composite-resource.deleteing")

	// Subscribe to the subject
	err = compositeResourcePendingSubject.Subscribe(cp.HandlePending)
	err = compositeResourceProvisioningSubject.Subscribe(cp.HandleProvisioning)
	err = compositeResourceDeletingSubject.Subscribe(cp.HandleDeleting)
	if err != nil {
		log.Logger.Fatal("Failed to subscribe message-queue", "error", err)
	}
}
