package handlers

import (
	"errors"
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

// BondErrorHandler is a decorator that is responsible for decorating the error handling functionality
func BondErrorHandler(apiFunc response.HttpHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(w, r); err != nil {
			asBond := domain.BondError(0)

			if errors.As(err, &asBond) {
				message := response.Response{
					Code:    asBond.Error(),
					Message: err.Error(),
				}

				switch asBond {
				case domain.InvalidBondID,
					domain.InvalidBondName,
					domain.InvalidBondQuantitySale,
					domain.InvalidBondSalesPrice,
					domain.DuplicateBond,
					domain.InvalidBondBought:
					_ = response.JSON(w, http.StatusBadRequest, message)
					return

				case domain.BondNotFound:
					_ = response.JSON(w, http.StatusNotFound, message)
					return
				default:
					_ = response.JSON(w, http.StatusInternalServerError, message)
					return
				}
			}

			asUser := domain.UserError(0)
			if errors.As(err, &asUser) {
				message := response.Response{
					Code:    asUser.Error(),
					Message: err.Error(),
				}

				switch asUser {
				case domain.InvalidUserID:
					_ = response.JSON(w, http.StatusBadRequest, message)
					return
				default:
					_ = response.JSON(w, http.StatusInternalServerError, message)
					return
				}
			}

			asFilter := filter.FilterError(0)
			if errors.As(err, &asFilter) {
				message := response.Response{
					Code:    asFilter.Error(),
					Message: err.Error(),
				}

				switch asFilter {
				case filter.InvalidFilterLimit, filter.InvalidFilterPage:
					_ = response.JSON(w, http.StatusUnprocessableEntity, message)
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
