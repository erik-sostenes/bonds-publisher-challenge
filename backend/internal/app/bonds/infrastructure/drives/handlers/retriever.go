package handlers

import (
	"net/http"
	"strconv"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// GetBondsPerUserHandler retrieve bonds posted by a specific user with pagination
func GetBondsPerUserHandler(retriever ports.UserBondsRetriever) response.HttpHandlerFunc {
	if retriever == nil {
		panic("missing User Bonds Retriever dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) (err error) {
		ctx := r.Context()

		bondCurrentOwnerId, err := domain.BondCurrentOwnerId(r.URL.Query().Get("current_owner_id")).Validate()
		if err != nil {
			return
		}

		page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 34)
		if err != nil {
			return response.JSON(w, http.StatusBadRequest, response.Response{
				Message: err.Error(),
			})
		}

		limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
		if err != nil {
			return response.JSON(w, http.StatusBadRequest, response.Response{
				Message: err.Error(),
			})
		}

		fltr, err := filter.NewFilter(page, limit)
		if err != nil {
			return
		}

		bonds, err := retriever.Retrieve(ctx, bondCurrentOwnerId, fltr)
		if err != nil {
			return
		}

		return response.JSON(w, http.StatusOK, toRequest(bonds))
	}
}
