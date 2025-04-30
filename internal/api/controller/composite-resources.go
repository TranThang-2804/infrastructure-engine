package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
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
	w.Header().Set("Content-Type", "application/json")
	var compositeResource domain.CompositeResource
	err := json.NewDecoder(r.Body).Decode(&compositeResource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Logger.Error("Error parsing composite resource config", "error", err.Error(), "compositeResourceConfig", compositeResource)
		return
	}
	compositeResource, err = rc.CompositeResourceUseCase.Create(r.Context(), compositeResource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Logger.Error("Error creating resource config", "error", err.Error(), "compositeResourceConfig", compositeResource)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(compositeResource)
}
