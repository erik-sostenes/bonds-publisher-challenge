package main

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// HealthCheck http handler that checks the status of the server
func HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = response.JSON(w, http.StatusOK, nil)
	}
}
