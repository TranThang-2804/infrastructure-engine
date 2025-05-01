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

func NewCompositeResourceRouter(env *bootstrap.Env, gitStore git.GitStore, timeout time.Duration, router chi.Router) {
	cr := repository.NewCompositeResourceRepository(gitStore)
  br := repository.NewBluePrintRepository(gitStore)
  bu := usecase.NewBluePrintUsecase(br, timeout)
	cp := &controller.CompositeResourceController{
		CompositeResourceUseCase: usecase.NewCompositeResourceUsecase(cr, bu, timeout),
		Env:                      env,
	}
	router.Get("/composite", cp.GetAll)
	router.Post("/composite", cp.Create)
}
