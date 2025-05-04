package controller

import (
	"net/http"
)

type HealthcheckController struct{}

func (hc *HealthcheckController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
