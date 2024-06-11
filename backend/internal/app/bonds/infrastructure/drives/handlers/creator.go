package handlers

import (
	"fmt"
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

		fmt.Println(bondRequest)

		return response.JSON(w, http.StatusCreated, nil)
	}
}
