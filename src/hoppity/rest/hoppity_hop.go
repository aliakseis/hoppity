package rest

import (
	"hoppity/api"
	"hoppity/models"
	"hoppity/util"
	"net/http"
	"fmt"
	"log"
)

// A new hoppity hop in POST body

func HoppityHop(w http.ResponseWriter, r *http.Request) {
  // swagger:operation POST /hoppity service hoppity
  //
  // Generates Hoppity Hop record.
  //
  // Returns Hoppity Hop status.
  //
  // ---
  // consumes:
  //  - application/json
  // parameters:
  //  - name: SequenceId
  //    in: body
  //    description: sequence ID
  //    required: true
  //    schema:
  //      "$ref": "#/definitions/SequenceId"
  // responses:
  //   '200':
  //     $ref: "#/responses/status"
  //   '500':
  //     $ref: "#/responses/status"
	var insertStruct models.SequenceId
	response := models.Status{ Status: "success"}

	err := util.ParseJsonObj(r, &insertStruct)
	if err != nil {
		response.Status="error"
		response.StatusMessage = fmt.Sprintf("Failed to get parse JSON object: %v", err)
		log.Print(response.StatusMessage)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}
	response, err = api.HoppityHop(insertStruct)
	if err != nil {
		log.Print(response)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}

	util.RespondWithJson(w, http.StatusOK, response )
}
