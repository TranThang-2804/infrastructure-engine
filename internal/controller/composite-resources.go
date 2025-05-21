package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

type CompositeResourceController struct {
	CompositeResourceUseCase domain.CompositeResourceUsecase
}

func NewCompositeResourceController(compositeResourceUseCase domain.CompositeResourceUsecase) *CompositeResourceController {
	return &CompositeResourceController{
		CompositeResourceUseCase: compositeResourceUseCase,
	}
}

func (rc *CompositeResourceController) GetAll(w http.ResponseWriter, r *http.Request) {
	logger := log.BaseLogger.FromCtx(r.Context()).WithFields("controller", utils.GetStructName(rc))
	ctx := logger.WithCtx(r.Context())

	compositeResources, err := rc.CompositeResourceUseCase.GetAll(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("Error getting all composite resources", "error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(compositeResources)
	logger.Info("Request handled Successful")
}

func (rc *CompositeResourceController) Create(w http.ResponseWriter, r *http.Request) {
	logger := log.BaseLogger.FromCtx(r.Context()).WithFields("controller", utils.GetStructName(rc))
	ctx := logger.WithCtx(r.Context())

	var request domain.CreateCompositeResourceRequest

	// Parse request body
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, utils.JsonError(err.Error()), http.StatusBadRequest)
		logger.Error("Error parsing body of creating resource api", "error", err.Error(), "compositeResourceConfig", request)
		return
	}

	// Validate request
	err = utils.ValidateStruct(request)
	if err != nil {
		http.Error(w, utils.JsonError(err.Error()), http.StatusBadRequest)
		logger.Error("Error validating request", "error", err.Error(), "compositeResourceConfig", request)
		return
	}

	compositeResource, err := rc.CompositeResourceUseCase.Create(ctx, request)
	if err != nil {
		http.Error(w, utils.JsonError(err.Error()), http.StatusInternalServerError)
		logger.Error("Error creating resource config", "error", err.Error(), "compositeResourceConfig", compositeResource)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(compositeResource)
	logger.Info("Request handled Successful")
}

func (rc *CompositeResourceController) Delete(w http.ResponseWriter, r *http.Request) {
}

func (rc *CompositeResourceController) Update(w http.ResponseWriter, r *http.Request) {
}

func (rc *CompositeResourceController) HandlePending(message string) error {
	return nil
}

func (rc *CompositeResourceController) HandleProvisioning(message string) error {
	return nil
}

func (rc *CompositeResourceController) HandleDeleting(message string) error {
	return nil
}
