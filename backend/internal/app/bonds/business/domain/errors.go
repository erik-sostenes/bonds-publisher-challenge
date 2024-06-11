package domain

import (
	"strconv"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
)

type (
	// UserError represents an error type for user-related errors
	UserError = domain.UserError
)

// InvalidUserID represents an error for an invalid user ID
const InvalidUserID = domain.InvalidUserID

const (
	// errors represents an error related to bond validation or business rules violation
	InvalidBondID BondError = iota + 1
	InvalidBondName
	InvalidBondQuantitySale
	InvalidBondSalesPrice
	DuplicateBond
)

// BondError represents an error type for bond-related errors
type BondError uint16

// BondError implements the error interface
func (b BondError) Error() string {
	return "bond: " + strconv.FormatUint(uint64(b), 10)
}
