package main

import (
	"context"
	"log"
	"merlin/controllers"
	"merlin/daos"
	merlin_logger "merlin/log"
	"merlin/routers"
	"merlin/usecases"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	merlin_logger.Initialise()

	allDaos := daos.Initialise(context.TODO())
	allUsecases := usecases.Initialize(allDaos)
	allControllers := controllers.Initialize(allUsecases)

	mainRouter := mux.NewRouter()
	routers.Initialize(mainRouter, allControllers)
	// n := middlewares.Initialize(mainRouter)

	// Start the server on port 2221.
	port := ":2221"
	log.Printf("Server is running on port %s\n", port)
	http.Handle("/", mainRouter)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
