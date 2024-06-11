package domain

import "strconv"

const (
	// InvalidUserID represents an error for an invalid user ID
	InvalidUserID UserError = iota + 1
)

// UserError represents an error type for user-related errors
type UserError uint16

// UserError implements the error interface
func (b UserError) Error() string {
	return "user: " + strconv.FormatUint(uint64(b), 10)
}
