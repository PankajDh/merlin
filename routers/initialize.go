package routers

import (
	"merlin/controllers"
	"reflect"

	"github.com/gorilla/mux"
)

type Routers struct {
	HealthRouter IRoute
}

func Initialize(mainRouter *mux.Router, allControllers *controllers.Controllers) *Routers {
	routers := &Routers{
		HealthRouter: NewHealthRouter(),
	}
	routers.registerRoutes(mainRouter, allControllers)
	return routers
}

// this function will create subroutes and registers endpoints
func (routers *Routers) registerRoutes(mainRouter *mux.Router, allControllers *controllers.Controllers) {
	t := reflect.Indirect(reflect.ValueOf(routers))

	for i := 0; i < t.NumField(); i++ {
		eachRouter := t.Field(i)
		basePath := eachRouter.Elem().FieldByName("BasePath").String()

		subRouter := mux.NewRouter().PathPrefix(basePath).Subrouter()

		addPathMethod := eachRouter.MethodByName("AddPaths")
		addPathMethod.Call([]reflect.Value{reflect.ValueOf(subRouter), reflect.ValueOf(allControllers)})

		mainRouter.PathPrefix(basePath).Handler(subRouter)
	}
}
