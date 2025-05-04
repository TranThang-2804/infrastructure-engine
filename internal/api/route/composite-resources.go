package route

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/go-chi/chi/v5"
)

func NewCompositeResourceRouter(router chi.Router, cp *controller.CompositeResourceController) {
	router.Get("/composite", cp.GetAll)
	router.Post("/composite", cp.Create)
}
