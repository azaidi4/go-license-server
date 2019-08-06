package api

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Simple Licensing Server")
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}
