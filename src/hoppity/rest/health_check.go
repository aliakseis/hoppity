package rest

import (
	"hoppity/api"
	"hoppity/util"
	"log"
	"net/http"
)

// Health Check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /health-check utility HealthCheckHandler
	// Performs application health check.
	// Returns 200 (OK) if health check succeeds, otherwise Internal Server Error (500) will be returned.
	// responses:
	//  200: status
	//  500: status
	response, err := api.HealthCheck(r.Context())
	if err != nil {
		log.Print(response)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}
	util.RespondWithJson(w, http.StatusOK, response)
}
