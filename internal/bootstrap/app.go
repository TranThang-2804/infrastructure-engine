package bootstrap

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/mq"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
)

type Application struct {
	// System Prerequisites
	infraPipeline InfraPipeline

	// Infrastructure Layer
	gitStore                   git.GitStore
	compositeResourcePublisher domain.CompositeResourceEventPublisher
	compositeResourceConsumer  domain.CompositeResourceEventConsumer
	mi                         mq.MessageQueue

	// Rrepository Layer
	compositeResourceRepository domain.CompositeResourceRepository
	bluePrintRepository         domain.BluePrintRepository
	iacTemplateRepository       domain.IacTemplateRepository
	iacPipelineRepository       domain.IacPipelineRepository

	// Usecase/Service Layer
	compositeResourceUsecase domain.CompositeResourceUsecase
	bluePrintUsecase         domain.BluePrintUsecase
	iacTemplateUsecase       domain.IacTemplateUsecase
	iacPipelineUsecase       domain.IacPipelineUsecase

	// Controller/Handler Layer
	CompositeResourceController *controller.CompositeResourceController
	BluePrintController         *controller.BluePrintController
	IacTemplateController       *controller.IacTemplateController
	HealthController            *controller.HealthController
}

func App() Application {
	mqSubjectList := []string{
		string(constant.ToPending),
		string(constant.ToProvisioning),
		string(constant.ToDeleting),
	}

	// Initiate the application
	app := &Application{}

	// Setting up infrastructure instances
	app.gitStore = NewGitHubStore()

	// Setting up required repositories pre-requisites
	app.infraPipeline = NewInfraPipeline(app.gitStore)
	app.infraPipeline.SettingInfraPipeline()

	// Create a mq infra type of NATS connection
	mi, err := mq.NewNatsMQ(env.Env.NATS_URL, mqSubjectList)
	if err != nil {
		log.BaseLogger.Fatal("Failed to connect to NATS", "error", err)
	}
	app.mi = mi
	app.compositeResourcePublisher = mq.NewCompositeResourcePublisher(app.mi)

	// Setting up the repositories
	app.compositeResourceRepository = repository.NewCompositeResourceRepository(app.gitStore)
	app.bluePrintRepository = repository.NewBluePrintRepository(app.gitStore)
	app.iacTemplateRepository = repository.NewIacTemplateRepository(app.gitStore)
	app.iacPipelineRepository = repository.NewIacPipelineRepository(app.gitStore)

	// Setting up the usecases
	app.bluePrintUsecase = usecase.NewBluePrintUsecase(app.bluePrintRepository)
	app.iacTemplateUsecase = usecase.NewIacTemplateUsecase(app.iacTemplateRepository)
	app.iacPipelineUsecase = usecase.NewIacPipelineUsecase(app.iacPipelineRepository)
	app.compositeResourceUsecase = usecase.NewCompositeResourceUsecase(app.compositeResourceRepository, app.compositeResourcePublisher, app.bluePrintUsecase, app.iacPipelineUsecase)

	// Setting up the controllers
	app.HealthController = controller.NewHealthController()
	app.BluePrintController = controller.NewBluePrintController(app.bluePrintUsecase)
	app.CompositeResourceController = controller.NewCompositeResourceController(app.compositeResourceUsecase)
	app.IacTemplateController = controller.NewIacTemplateController(app.iacTemplateUsecase)

	// Setting up the consumer
	app.compositeResourceConsumer = mq.NewCompositeResourceConsumer(app.mi, app.compositeResourceUsecase)
	app.compositeResourceConsumer.StartConsumer()

	return *app
}

func (app *Application) CloseDBConnection() {
	// Close the message queue connection
	app.mi.Close()
}
