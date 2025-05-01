package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
	"github.com/go-playground/validator/v10"
)

type CompositeResourceController struct {
	CompositeResourceUseCase domain.CompositeResourceUsecase
	Env                      *bootstrap.Env
}

func (rc *CompositeResourceController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	compositeResources, err := rc.CompositeResourceUseCase.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Logger.Error("Error getting all composite resources", "error", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(compositeResources)
}

func (rc *CompositeResourceController) Create(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateCompositeResourceRequest

  // Parse request body
  decoder := json.NewDecoder(r.Body)
  decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, utils.JsonError(err.Error()), http.StatusBadRequest)
		log.Logger.Error("Error parsing body of creating resource api", "error", err.Error(), "compositeResourceConfig", request)
		return
	}

  // Validate request
  validate := validator.New()
	err = validate.Struct(request)
  if err != nil {
    http.Error(w, utils.JsonError(err.Error()), http.StatusBadRequest)
    log.Logger.Error("Error validating request", "error", err.Error(), "compositeResourceConfig", request)
    return
  }

	compositeResource, err := rc.CompositeResourceUseCase.Create(r.Context(), request)
	if err != nil {
		http.Error(w, utils.JsonError(err.Error()), http.StatusInternalServerError)
		log.Logger.Error("Error creating resource config", "error", err.Error(), "compositeResourceConfig", compositeResource)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(compositeResource)
}
