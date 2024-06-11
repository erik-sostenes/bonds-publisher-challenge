package handlers

import (
	"net/http"
	"strconv"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// GetBondsPerUserHandler retrieve bonds posted by a specific user with pagination
func GetBondsPerUserHandler() response.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		_, err = domain.BondCurrentOwnerId(r.URL.Query().Get("current_owner_id")).Validate()
		if err != nil {
			return
		}

		_, err = strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
		if err != nil {
			return response.JSON(w, http.StatusBadRequest, response.Response{
				Message: err.Error(),
			})
		}

		_, err = strconv.ParseUint(r.URL.Query().Get("page"), 10, 34)
		if err != nil {
			return response.JSON(w, http.StatusBadRequest, response.Response{
				Message: err.Error(),
			})
		}

		return response.JSON(w, http.StatusOK, nil)
	}
}
