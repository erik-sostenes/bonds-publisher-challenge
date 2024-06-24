package handlers

import (
	"errors"
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// UserErrorHandler is a decorator that is responsible for decorating the error handling functionality
func UserErrorHandler(apiFunc response.HttpHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {
			asUser := domain.UserError(0)

			if errors.As(err, &asUser) {
				message := response.Response{
					Code:    asUser.Error(),
					Message: err.Error(),
				}

				switch asUser {
				case domain.InvalidUserID,
					domain.InvalidUserName,
					domain.InvalidUserPassword,
					domain.DuplicateUser:
					_ = response.JSON(w, http.StatusBadRequest, message)
					return
				case domain.InvalidToken, domain.PasswordDoesNotMatch:
					_ = response.JSON(w, http.StatusUnauthorized, message)
					return
				case domain.UserNotFound:
					_ = response.JSON(w, http.StatusNotFound, message)
					return

				default:
					_ = response.JSON(w, http.StatusInternalServerError, message)
					return
				}
			}

			asRole := domain.RoleError(0)
			if errors.As(err, &asRole) {
				message := response.Response{
					Code:    asRole.Error(),
					Message: err.Error(),
				}

				switch asRole {
				case domain.InvalidRoleType,
					domain.InvalidRoleTypeID:
					_ = response.JSON(w, http.StatusBadRequest, message)
					return
				default:
					_ = response.JSON(w, http.StatusInternalServerError, message)
					return

				}
			}

			_ = response.JSON(w, http.StatusInternalServerError, "an error has occurred")
			return
		}
	}
}
