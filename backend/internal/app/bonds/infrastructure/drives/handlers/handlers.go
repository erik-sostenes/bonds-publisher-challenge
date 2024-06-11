package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	md "github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/middlewares"
)

// BondHandler configures bond end-points
func BondHandler(bondCtr ports.BondCreator, mux *http.ServeMux) {
	mux.HandleFunc(
		"POST /api/v1/bonds/create",
		md.Recovery(md.Logger(md.CORS(BondErrorHandler(PostBondHandler(bondCtr))))),
	)
}
