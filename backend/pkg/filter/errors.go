package filter

import "strconv"

const (
	// errors represents an error related to the validation of the filter business rules
	InvalidFilterPage FilterError = iota + 1
	InvalidFilterLimit
)

// FilterError represents an error type for filter-related errors
type FilterError uint16

// FilterError implements the error interface
func (f FilterError) Error() string {
	return "filter: " + strconv.FormatUint(uint64(f), 10)
}
