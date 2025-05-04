package route

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/controller"
	"github.com/go-chi/chi/v5"
)

func NewIacTemplateRouter(router chi.Router, ip *controller.IacTemplateController) {
	router.Get("/iac-template", ip.GetAll)
}
