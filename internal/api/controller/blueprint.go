package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type BluePrintController struct {
	BluePrintUsecase domain.BluePrintUsecase
	Env              *bootstrap.Env
}

func (bp *BluePrintController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
  emptyBlueprint := []domain.BluePrint{}
	json.NewEncoder(w).Encode(emptyBlueprint)
}
