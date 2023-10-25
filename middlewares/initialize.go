package middlewares

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Initialize(mainRouter *mux.Router) *negroni.Negroni {
	n := negroni.New()
	return n
}
