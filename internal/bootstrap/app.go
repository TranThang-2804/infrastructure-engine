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
	iaCTemplateRepository       domain.IacTemplateRepository

	// Usecase/Service Layer
	compositeResourceUsecase domain.CompositeResourceUsecase
	bluePrintUsecase         domain.BluePrintUsecase
	iacTemplateUsecase       domain.IacTemplateUsecase

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
	// Create a mq infra type of NATS connection
	mi, err := mq.NewNatsMQ(env.Env.NATS_URL, mqSubjectList)
	if err != nil {
		log.Logger.Fatal("Failed to connect to NATS", "error", err)
	}
	app.mi = mi
	app.compositeResourcePublisher = mq.NewCompositeResourcePublisher(app.mi)

	// Setting up the repositories
	app.compositeResourceRepository = repository.NewCompositeResourceRepository(app.gitStore)
	app.bluePrintRepository = repository.NewBluePrintRepository(app.gitStore)
	app.iaCTemplateRepository = repository.NewIacTemplateRepository(app.gitStore)

	// Setting up the usecases
	app.bluePrintUsecase = usecase.NewBluePrintUsecase(app.bluePrintRepository)
	app.iacTemplateUsecase = usecase.NewIacTemplateUsecase(app.iaCTemplateRepository)
	app.compositeResourceUsecase = usecase.NewCompositeResourceUsecase(app.compositeResourceRepository, app.compositeResourcePublisher, app.bluePrintUsecase)

	// Setting up the controllers
	app.HealthController = controller.NewHealthController()
	app.BluePrintController = controller.NewBluePrintController(app.bluePrintUsecase)
	app.CompositeResourceController = controller.NewCompositeResourceController(app.compositeResourceUsecase)
	app.IacTemplateController = controller.NewIacTemplateController(app.iacTemplateUsecase)

	// Setting up the consumer
	app.compositeResourceConsumer = mq.NewCompositeResourceConsumer(app.mi, app.compositeResourceUsecase)
	app.compositeResourceConsumer.StartConsumer()

	// Setting up required repositories pre-requisites
	app.infraPipeline = NewInfraPipeline(app.gitStore)
	app.infraPipeline.SettingInfraPipeline()
	return *app
}

func (app *Application) CloseDBConnection() {
	// Close the message queue connection
	app.mi.Close()
}
