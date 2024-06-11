package filter

import "fmt"

const (
	MinLengthFilterPage = 1
	LengthFilterLimit   = 25
)

// FilterPage represents the page number for pagination
type FilterPage uint64

func (o FilterPage) Validate() (*FilterPage, error) {
	if o < MinLengthFilterPage {
		return nil, fmt.Errorf("%w = the page value is less than %v", InvalidFilterPage, MinLengthFilterPage)
	}

	return &o, nil
}

func (o FilterPage) Page() uint64 {
	return uint64(o)
}

// FilterLimit represents the limit of items per page
type FilterLimit uint64

func (o FilterLimit) Validate() (*FilterLimit, error) {
	if o != LengthFilterLimit {
		return nil, fmt.Errorf("%w = the limit per page is different from %v", InvalidFilterLimit, LengthFilterLimit)
	}

	return &o, nil
}

func (o FilterLimit) Limit() uint64 {
	return uint64(o)
}

// FilterSize represents the size of the order filter
type FilterSize uint8

func (o FilterSize) Validate() (*FilterSize, error) {
	return &o, nil
}

func (o FilterSize) Size() uint64 {
	return uint64(o)
}

// Filter holds the pagination details including page number, limit, and size
type Filter struct {
	orderFilterPage  FilterPage
	orderFilterLimit FilterLimit
	orderFilterSize  FilterSize
}

func NewFilter(page, limit uint64) (*Filter, error) {
	pageVO, err := FilterPage(page).Validate()
	if err != nil {
		return &Filter{}, err
	}

	limitVO, err := FilterLimit(limit).Validate()
	if err != nil {
		return &Filter{}, err
	}
	return &Filter{
		orderFilterPage:  *pageVO,
		orderFilterLimit: *limitVO,
	}, nil
}

// Stop calculates and returns the stop index for pagination.
func (o *Filter) Stop() uint64 {
	return o.orderFilterPage.Page() * o.orderFilterLimit.Limit()
}

// Start calculates and returns the start index for pagination.
func (o *Filter) Start() uint64 {
	return (o.orderFilterPage.Page() - 1) * o.orderFilterLimit.Limit()
}

// SetSize sets the size value for the order filter.
func (o *Filter) SetSize(size uint8) {
	o.orderFilterSize = FilterSize(size)
}

// Size returns the size value of the order filter.
func (o *Filter) Size() uint64 {
	return o.orderFilterSize.Size()
}
