package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type BluePrintController struct {
	BluePrintUsecase domain.BluePrintUsecase
}

func NewBluePrintController(bluePrintUsecase domain.BluePrintUsecase) *BluePrintController {
	return &BluePrintController{
		BluePrintUsecase: bluePrintUsecase,
	}
}

func (bc *BluePrintController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bluePrints, err := bc.BluePrintUsecase.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bluePrints)
}
