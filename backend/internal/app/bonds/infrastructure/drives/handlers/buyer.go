package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// PutBondBuyerHandler http handler that receives the http request to buy a bond
func PutBondBuyerHandler(buyer ports.BondBuyer) response.HttpHandlerFunc {
	if buyer == nil {
		panic("missing Bond Buyer dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) (err error) {
		ctx := r.Context()

		pathVBondID := r.PathValue("bond_id")
		bondID, err := domain.BondID(pathVBondID).Validate()
		if err != nil {
			return
		}

		pathVBuyerUserId := r.PathValue("buyer_user_id")
		buyerUserID, err := domain.BondCurrentOwnerId(pathVBuyerUserId).Validate()
		if err != nil {
			return
		}

		if err = buyer.Buy(ctx, bondID, buyerUserID); err != nil {
			return
		}

		return response.JSON(w, http.StatusOK, nil)
	}
}
