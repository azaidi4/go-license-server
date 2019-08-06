package api

import (
	"fmt"
	"license-server/env"
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

	r.NewRoute().Path("/check").
		Methods(http.MethodPost).
		HandlerFunc(checkHandler)

	return r
}

// StartServer spins up license-server
func StartServer() {
	r := router()

	server := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    env.Config.Server.URL,
	}

	fmt.Println("Listening to", env.Config.Server.URL)
	log.Fatal(server.ListenAndServe())
}
