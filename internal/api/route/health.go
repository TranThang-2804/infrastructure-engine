package route

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/api/controller"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/go-chi/chi/v5"
)

func NewHealthCheckRouter(env *bootstrap.Env, timeout time.Duration, router chi.Router) {
	hc := &controller.HealthcheckController{}
	router.Get("/health", hc.HealthCheck)
}
