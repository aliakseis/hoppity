package api

import (
	"context"
	"hoppity/models"
)

// Health Check
func HealthCheck(ctx context.Context) (payload interface{}, err error) {
	status := models.Status{Status: "success"}
	return status, nil
}
