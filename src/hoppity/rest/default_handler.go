package rest

import (
	"net/http"
	"log"
)

// Serve a static page for this service
func DefaultHandler(w http.ResponseWriter, _ *http.Request) {
	log.Print("In DefaultHandler")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write( []byte ("This is the HoppityHop service"))
}
