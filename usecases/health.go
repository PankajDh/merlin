package usecases

import (
	"merlin/models"

	"github.com/mackerelio/go-osstat/uptime"
	"github.com/pkg/errors"
)

type HealthUsecase struct {
}

func NewHealthUsecase() *HealthUsecase {
	return &HealthUsecase{}
}

func (healthUseCase *HealthUsecase) Check() (healhCheck models.HealthCheckResponse, err error) {
	serverUptime, err := uptime.Get()
	if err != nil {
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return healhCheck, errors.WithStack(err)
	}

	healhCheck = models.HealthCheckResponse{
		Uptime: serverUptime,
		Status: "OK",
	}

	return healhCheck, nil
}
