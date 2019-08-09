package api

import (
	"license-server/env"
	"license-server/utils/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()

	r.NewRoute().Path("/").
		Methods(http.MethodGet).
		HandlerFunc(indexHandler)

	r.NewRoute().Path("/validate").
		Methods(http.MethodPost).
		HandlerFunc(validateHandler)

	return r
}

// StartServer spins up license-server
func StartServer() {
	logger.Info.Println("Starting up License Server...")

	r := router()
	logging := handlers.LoggingHandler(os.Stdout, r)
	server := &http.Server{
		Handler: logging,
		Addr:    env.Config.Server.URL,
	}

	logger.Info.Println("Listening to", env.Config.Server.URL)
	log.Fatal(server.ListenAndServe())
}
