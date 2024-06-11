package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// PutBondBuyerHandler http handler that receives the http request to buy a bond
func PutBondBuyerHandler() response.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		pathVBondID := r.PathValue("bond_id")
		_, err = domain.BondID(pathVBondID).Validate()
		if err != nil {
			return
		}

		pathVBuyerUserId := r.PathValue("buyer_user_id")
		_, err = domain.BondCurrentOwnerId(pathVBuyerUserId).Validate()
		if err != nil {
			return
		}

		return response.JSON(w, http.StatusOK, nil)
	}
}
