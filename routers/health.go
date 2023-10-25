package routers

import (
	"merlin/controllers"

	"github.com/gorilla/mux"
)

type HealthRouter struct {
	BasePath string
}

func NewHealthRouter() IRoute {
	return HealthRouter{
		BasePath: "/api/v1/health",
	}
}

func (healthRouter HealthRouter) AddPaths(localRouter *mux.Router, allControllers *controllers.Controllers) {
	healthController := allControllers.HealthController

	localRouter.HandleFunc("", healthController.Check).Methods("GET")
}
