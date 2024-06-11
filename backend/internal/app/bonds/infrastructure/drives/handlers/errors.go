package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// BondErrorHandler is a decorator that is responsible for decorating the error handling functionality
func BondErrorHandler(apiFunc response.HttpHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {
			return
		}
	}
}
