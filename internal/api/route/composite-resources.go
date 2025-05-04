package route

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewCompositeResourceRouter(gitStore git.GitStore, timeout time.Duration, router chi.Router) {
	cr := repository.NewCompositeResourceRepository(gitStore)
	br := repository.NewBluePrintRepository(gitStore)
	bu := usecase.NewBluePrintUsecase(br, timeout)
	cp := &controller.CompositeResourceController{
		CompositeResourceUseCase: usecase.NewCompositeResourceUsecase(cr, bu, timeout),
	}
	router.Get("/composite", cp.GetAll)
	router.Post("/composite", cp.Create)
}
