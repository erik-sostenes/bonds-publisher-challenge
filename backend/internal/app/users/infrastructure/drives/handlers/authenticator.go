package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func GetAuthenticator(authorizer ports.UserAuthorizer) response.HttpHandlerFunc {
	if authorizer == nil {
		panic("missing 'User Authorizer' dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) (err error) {
		ctx := r.Context()

		username, err := domain.UserName(r.URL.Query().Get("username")).Validate()
		if err != nil {
			return
		}

		userPassword, err := domain.UserPassword(r.URL.Query().Get("user_password")).Validate()
		if err != nil {
			return
		}

		token, err := authorizer.Authorize(ctx, username, userPassword)
		if err != nil {
			return
		}

		return response.JSON(w, http.StatusOK, map[string]any{"token": token})
	}
}
