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

func NewIacTemplateRouter(env *bootstrap.EnvConfig, gitStore git.GitStore, timeout time.Duration, router chi.Router) {
	ir := repository.NewIacTemplateRepository(gitStore)
	ip := &controller.IacTemplateController{
		IacTemplateUsecase: usecase.NewIacTemplateUsecase(ir, timeout),
		Env:                env,
	}
	router.Get("/iac-template", ip.GetAll)
}
