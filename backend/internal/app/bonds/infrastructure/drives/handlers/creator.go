package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// PostBondHandler http handler that receives the http request to create a new bond
func PostBondHandler() response.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		var bondRequest BondRequest
		if err = response.Bind(w, r, &bondRequest); err != nil {
			return
		}

		return response.JSON(w, http.StatusCreated, nil)
	}
}
