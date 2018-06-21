package api

import (
	"hoppity/models"

	"log"
)

// A new hoppity hop in POST body
func HoppityHop(seqId models.SequenceId) (models.Status, error) {

	i := seqId.SequenceId

	log.Print("In HoppityHop ")

	status := models.Status{Status: "success"}

	if i % 15 == 0 {
		status.StatusMessage = "Hop"
	} else if i % 5 == 0 {
		status.StatusMessage = "Hophop"
	} else if i % 3 == 0 {
		status.StatusMessage = "Hoppity"
	}

	return status, nil
}
