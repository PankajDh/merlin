package controllers

import (
	"merlin/usecases"
)

type AWSController struct {
	AWSUsecase *usecases.AWSUsecase
}

func NewAWSController(awsUsecase *usecases.AWSUsecase) *AWSController {
	return &AWSController{
		AWSUsecase: awsUsecase,
	}
}
