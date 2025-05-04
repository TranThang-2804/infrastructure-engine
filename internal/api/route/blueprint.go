package route

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/go-chi/chi/v5"
)

func NewBluePrintRouter(router chi.Router, bp *controller.BluePrintController) {
	router.Get("/blueprint", bp.GetAll)
}
