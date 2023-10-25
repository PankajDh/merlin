package controllers

import "merlin/usecases"

type Controllers struct {
	HealthController *HealthController
	AWSController    *AWSController
}

func Initialize(allUseCases *usecases.Usecases) *Controllers {
	return &Controllers{
		HealthController: NewHealthController(allUseCases.HealthUsecase),
		AWSController:    NewAWSController(allUseCases.AWSUsecase),
	}
}
