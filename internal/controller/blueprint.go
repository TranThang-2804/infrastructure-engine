package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
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
	logger := log.BaseLogger.FromCtx(r.Context()).WithFields("controller", utils.GetStructName(bc))
	ctx := logger.WithCtx(r.Context())

	logger.Info("Processing health check request")

	bluePrints, err := bc.BluePrintUsecase.GetAll(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("Request handled failed", "error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bluePrints)
	logger.Info("Request handled Successful")
}
