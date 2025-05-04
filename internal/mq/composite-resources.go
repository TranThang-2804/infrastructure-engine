package mq

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/message-queue"
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
)

func SetupMQController(gitStore git.GitStore, timeout time.Duration) {
	mqSubjectList := []string{
		string(constant.ToPending),
		string(constant.ToProvisioning),
		string(constant.ToDeleting),
	}

	// Create a NATS connection
	mq, err := mqadapter.NewNatsMQ(env.Env.NATS_URL, mqSubjectList)
	if err != nil {
		log.Logger.Fatal("Failed to connect to NATS", "error", err)
	}

	cr := repository.NewCompositeResourceRepository(gitStore)
	br := repository.NewBluePrintRepository(gitStore)
	bu := usecase.NewBluePrintUsecase(br, timeout)
	cp := &controller.CompositeResourceController{
		CompositeResourceUseCase: usecase.NewCompositeResourceUsecase(cr, bu, timeout),
	}
}
