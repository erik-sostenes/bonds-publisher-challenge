package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	md "github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/middlewares"
)

// BondHandler configures bond end-points
func BondHandler(
	creator ports.BondCreator,
	buyer ports.BondBuyer,
	usrBondsRetriever ports.UserBondsRetriever,
	bondsRetriever ports.BondsRetriever,
	mux *http.ServeMux,
) {
	mux.HandleFunc(
		"POST /api/v1/bonds/create",
		md.Recovery(md.Logger(md.CORS(BondErrorHandler(PostBondHandler(creator))))),
	)
	mux.HandleFunc(
		"PUT /api/v1/bonds/buy/{bond_id}/{buyer_user_id}",
		md.Recovery(md.Logger(md.CORS(BondErrorHandler(PutBondBuyerHandler(buyer))))),
	)
	mux.HandleFunc(
		"GET /api/v1/bonds/user",
		md.Recovery(md.Logger(md.CORS(BondErrorHandler(GetBondsPerUserHandler(usrBondsRetriever))))),
	)
	mux.HandleFunc(
		"GET /api/v1/bonds/all",
		md.Recovery(md.Logger(md.CORS(BondErrorHandler(GetBondsHandler(bondsRetriever))))),
	)
}
