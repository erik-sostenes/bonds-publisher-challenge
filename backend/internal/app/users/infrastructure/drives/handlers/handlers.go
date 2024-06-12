package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	md "github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/middlewares"
)

func UserHandler(creator ports.UserCreator, mux *http.ServeMux) {
	mux.HandleFunc(
		"POST /api/v1/register",
		md.Recovery(md.Logger(md.CORS(UserErrorHandler(PostUserHandler(creator))))),
	)
}
