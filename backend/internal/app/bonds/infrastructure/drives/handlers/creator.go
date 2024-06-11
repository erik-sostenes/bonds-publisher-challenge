package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// PostBondHandler http handler that receives the http request to create a new bond
func PostBondHandler(creator ports.BondCreator) response.HttpHandlerFunc {
	if creator == nil {
		panic("missing Bond Creator dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) (err error) {
		ctx := r.Context()

		var bondRequest BondRequest
		if err = response.Bind(w, r, &bondRequest); err != nil {
			return
		}

		bond, err := bondRequest.toBusiness()
		if err != nil {
			return err
		}

		if err := creator.Create(ctx, bond); err != nil {
			return err
		}

		return response.JSON(w, http.StatusCreated, nil)
	}
}
