package handlers

import (
	"net/http"

	bondPorts "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	usrPorts "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	md "github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/middlewares"
)

// BondHandler configures bond end-points
func BondHandler(
	creator bondPorts.BondCreator,
	buyer bondPorts.BondBuyer,
	usrBondsRetriever bondPorts.UserBondsRetriever,
	bondsRetriever bondPorts.BondsRetriever,
	tokenValidator usrPorts.TokenValidator,
	mux *http.ServeMux,
) {
	mux.HandleFunc(
		"POST /api/v1/bonds/create",
		md.Recovery(md.Logger(md.IsAuthenticated(tokenValidator, BondErrorHandler(PostBondHandler(creator))))),
	)

	mux.HandleFunc(
		"PUT /api/v1/bonds/buy/{bond_id}/{buyer_user_id}",
		md.Recovery(md.Logger(BondErrorHandler(PutBondBuyerHandler(buyer)))),
	)

	mux.HandleFunc(
		"GET /api/v1/bonds/user",
		md.Recovery(md.Logger(BondErrorHandler(GetBondsPerUserHandler(usrBondsRetriever)))),
	)

	mux.HandleFunc(
		"GET /api/v1/bonds/all",
		md.Recovery(md.Logger(BondErrorHandler(GetBondsHandler(bondsRetriever)))),
	)
}
