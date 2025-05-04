package route

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/go-chi/chi/v5"
)

func NewHealthRouter(router chi.Router, hc *controller.HealthController) {
	router.Get("/health", hc.HealthCheck)
}
