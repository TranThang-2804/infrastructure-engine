package route

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/repository"
	"github.com/TranThang-2804/infrastructure-engine/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewBluePrintRouter(gitStore git.GitStore, timeout time.Duration, router chi.Router) {
	br := repository.NewBluePrintRepository(gitStore)
	bp := &controller.BluePrintController{
		BluePrintUsecase: usecase.NewBluePrintUsecase(br, timeout),
	}
	router.Get("/blueprint", bp.GetAll)
}
