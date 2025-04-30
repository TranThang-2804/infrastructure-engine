package route

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/api/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewBluePrintRouter(env *bootstrap.Env, timeout time.Duration, router chi.Router) {
	br := repository.NewBluePrintRepository()
	bp := &controller.BluePrintController{
		BluePrintUsecase: usecase.NewBluePrintUsecase(br, timeout),
		Env:              env,
	}
	router.Get("/blueprint", bp.GetAll)
}
