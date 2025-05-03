package route

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/api/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewBluePrintRouter(env *bootstrap.EnvConfig, gitStore git.GitStore, timeout time.Duration, router chi.Router) {
	br := repository.NewBluePrintRepository(gitStore)
	bp := &controller.BluePrintController{
		BluePrintUsecase: usecase.NewBluePrintUsecase(br, timeout),
		Env:              env,
	}
	router.Get("/blueprint", bp.GetAll)
}
