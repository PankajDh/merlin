package usecases

import "merlin/daos"

type Usecases struct {
	HealthUsecase *HealthUsecase
	AWSUsecase    *AWSUsecase
	Daos          *daos.Daos
}

func Initialize(daos *daos.Daos) *Usecases {
	return &Usecases{
		HealthUsecase: NewHealthUsecase(),
		AWSUsecase:    NewAWSUsecase(),
		Daos:          daos,
	}
}
