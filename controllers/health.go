package controllers

import (
	"encoding/json"
	"merlin/usecases"
	"net/http"
)

type HealthController struct {
	HealthUsecase *usecases.HealthUsecase
}

func NewHealthController(healthUsecase *usecases.HealthUsecase) *HealthController {
	return &HealthController{
		HealthUsecase: healthUsecase,
	}
}

func (healthController *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	responsePayload, err := healthController.HealthUsecase.Check()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jsonPayload, err := json.Marshal(responsePayload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonPayload))
}
