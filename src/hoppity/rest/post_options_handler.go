package rest

import (
	"net/http"
	"log"
)

// Swagger UI sends OPTIONS request method instead of POST
func PostOptionsHandler(w http.ResponseWriter, _ *http.Request) {
	log.Print("In PostOptionsHandler")

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "API-Key,accept, Content-Type")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Content-Length", "0")

	w.WriteHeader(http.StatusOK)
}
