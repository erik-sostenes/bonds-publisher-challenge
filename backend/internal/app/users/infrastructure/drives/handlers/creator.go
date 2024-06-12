package handlers

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// PostUserHandler http handler that receives the http request to create a new user
func PostUserHandler(creator ports.UserCreator) response.HttpHandlerFunc {
	if creator == nil {
		panic("missing User Creator dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) (err error) {
		ctx := r.Context()

		var userRequest UserRequest
		if err = response.Bind(w, r, &userRequest); err != nil {
			return
		}

		user, err := userRequest.toBusiness()
		if err != nil {
			return
		}

		if err = creator.Create(ctx, user); err != nil {
			return
		}

		return response.JSON(w, http.StatusCreated, nil)
	}
}
