package api

import (
	"hoppity/models"
	"testing"
)

func TestHoppityHop(t *testing.T) {

	status, err := HoppityHop(models.SequenceId{SequenceId: 15})
	if err != nil {
		t.Fail()
	}
	if status.Status != "success" {
		t.Fail()
	}
	if status.StatusMessage != "Hop" {
		t.Fail()
	}
}
