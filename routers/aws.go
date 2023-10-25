package routers

import (
	"merlin/controllers"

	"github.com/gorilla/mux"
)

type AWSRouter struct {
	BasePath string
}

func NewAWSRouter() IRoute {
	return AWSRouter{
		BasePath: "/api/v1/aws",
	}
}

func (awsRouter AWSRouter) AddPaths(localRouter *mux.Router, allControllers *controllers.Controllers) {
	// awsController := allControllers.AWSController

	// localRouter.HandleFunc("/connections", awsController.Check).Methods("GET")
}
