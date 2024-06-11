package handlers

import (
	"net/http"

	md "github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/middlewares"
)

// BondHandler configures bond end-points
func BondHandler(mux *http.ServeMux) {
	mux.HandleFunc(
		"POST /api/v1/bonds/create",
		md.Recovery(md.Logger(md.CORS(BondErrorHandler(PostBondHandler())))),
	)
}
