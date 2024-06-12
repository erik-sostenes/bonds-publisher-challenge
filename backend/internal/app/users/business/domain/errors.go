package domain

import "strconv"

const (
	// errors represents an error related to user validation of the filter business rules
	InvalidUserID UserError = iota + 1
	InvalidUserName
	InvalidUserPassword
	DuplicateUser
)

// UserError represents an error type for user-related errors
type UserError uint16

// UserError implements the error interface
func (b UserError) Error() string {
	return "user: " + strconv.FormatUint(uint64(b), 10)
}

const (
	// errors represents an error related to roles validation of the filter business rules
	InvalidRoleType RoleError = iota + 1
	InvalidRoleTypeID
)

// RoleError represents an error type for role-related errors
type RoleError uint16

// RoleError implements the error interface
func (r RoleError) Error() string {
	return "role: " + strconv.FormatUint(uint64(r), 10)
}
