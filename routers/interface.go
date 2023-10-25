package routers

import (
	"merlin/controllers"

	"github.com/gorilla/mux"
)

type BaseRouter struct {
	BasePath string
}

type IRoute interface {
	AddPaths(localRouter *mux.Router, allControllers *controllers.Controllers)
}
