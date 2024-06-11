package domain

import (
	"fmt"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/google/uuid"
)

const (
	MinLengthBondName = 3
	MaxLengthBondName = 40

	MinBondQuantitySale = 1
	MaxBondQuantitySale = 10000

	MinBondSalesPrice = 0.0000
	MaxBondSalesPrice = 100000000.0000
)

// BondID represents the unique identifier for a bond
type BondID string

// Validate checks if the BondID is a valid UUID format
func (b BondID) Validate() (*BondID, error) {
	_, err := uuid.Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("%w = %s", InvalidBondID, err.Error())
	}

	return &b, nil
}

// ID returns the string representation of the BondID
func (b BondID) ID() string {
	return string(b)
}

// BondName represents the name of a bond
type BondName string

// Validate ensures that the BondName length is within the specified range
func (b BondName) Validate() (*BondName, error) {
	if len(b) < MinLengthBondName || len(b) > MaxLengthBondName {
		return nil, fmt.Errorf("%w = %s", InvalidBondName, "the size of the bond name is not within range")
	}

	return &b, nil
}

// Name returns the string representation of the BondName
func (b BondName) Name() string {
	return string(b)
}

// BondQuantitySale represents the quantity of bonds available for sale
type BondQuantitySale int

// Validate ensures that the BondQuantitySale is within the specified range
func (b BondQuantitySale) Validate() (*BondQuantitySale, error) {
	if b < MinBondQuantitySale || b > MaxBondQuantitySale {
		return nil, fmt.Errorf("%w = %s", InvalidBondQuantitySale, "the bond quantity sale is not within range")
	}
	return &b, nil
}

// QuantitySale returns the integer representation of the BondQuantitySale
func (b BondQuantitySale) QuantitySale() int {
	return int(b)
}

// BondSalesPrice represents the sales price of bonds
type BondSalesPrice float64

// Validate ensures that the BondSalesPrice is within the specified range
func (b BondSalesPrice) Validate() (*BondSalesPrice, error) {
	if b < MinBondSalesPrice || b > MaxBondSalesPrice {
		return nil, fmt.Errorf("%w = %s", InvalidBondSalesPrice, "the bond sales price is not within range")
	}

	return &b, nil
}

// SalesPrice returns the float64 representation of the BondSalesPrice
func (b BondSalesPrice) SalesPrice() float64 {
	return float64(b)
}

// BondIsBought represents whether a bond has been bought
type BondIsBought bool

// Validate always returns nil since no validation is needed
func (b BondIsBought) Validate() (*BondIsBought, error) {
	return &b, nil
}

// BondIsBought returns the boolean representation of the BondIsBought
func (b BondIsBought) BondIsBought() bool {
	return bool(b)
}

type (
	// BondCreatorUserID represents the unique identifier of the user who created the bond
	BondCreatorUserID = domain.UserID
	// BondCurrentOwnerId represents the unique identifier of the current owner of the bond
	BondCurrentOwnerId = domain.UserID

	// Bond represents the Object Domain of our business
	Bond struct {
		bondID             BondID
		bondName           BondName
		bondQuantitySale   BondQuantitySale
		bondSalesPrice     BondSalesPrice
		bondIsBought       BondIsBought
		bondCreatorUserId  BondCreatorUserID
		bondCurrentOwnerId BondCurrentOwnerId
	}
)

// NewBond creates a new Bond instance with the provided parameters
func NewBond(
	bondId,
	bondName,
	bondCreatorUserId,
	bondCurrentOwnerId string,
	bondIsBought bool,
	bondQuantitySale int64,
	bondSalesPrice float64,
) (*Bond, error) {
	bondIdVO, err := BondID(bondId).Validate()
	if err != nil {
		return nil, err
	}

	bondNameVO, err := BondName(bondName).Validate()
	if err != nil {
		return nil, err
	}

	bondQuantitySaleVO, err := BondQuantitySale(bondQuantitySale).Validate()
	if err != nil {
		return nil, err
	}

	bondIsBoughtVO, err := BondIsBought(bondIsBought).Validate()
	if err != nil {
		return nil, err
	}

	bondSalesPriceVO, err := BondSalesPrice(bondSalesPrice).Validate()
	if err != nil {
		return nil, err
	}

	bondCreatorUserIdVO, err := BondCreatorUserID(bondCreatorUserId).Validate()
	if err != nil {
		return nil, err
	}

	bondCurrentOwnerIdVO, err := BondCurrentOwnerId(bondCurrentOwnerId).Validate()
	if err != nil {
		return nil, err
	}

	return &Bond{
		bondID:             *bondIdVO,
		bondName:           *bondNameVO,
		bondQuantitySale:   *bondQuantitySaleVO,
		bondSalesPrice:     *bondSalesPriceVO,
		bondIsBought:       *bondIsBoughtVO,
		bondCreatorUserId:  *bondCreatorUserIdVO,
		bondCurrentOwnerId: *bondCurrentOwnerIdVO,
	}, nil
}

// ID returns the ID of the Bond
func (b Bond) ID() string {
	return b.bondID.ID()
}

// Name returns the name of the Bond
func (b Bond) Name() string {
	return b.bondName.Name()
}

// QuantitySale returns the quantity of bonds available for sale
func (b Bond) QuantitySale() int {
	return b.bondQuantitySale.QuantitySale()
}

// SalesPrice returns the sales price of the Bond
func (b Bond) SalesPrice() float64 {
	return b.bondSalesPrice.SalesPrice()
}

// IsBought returns whether the Bond has been bought
func (b Bond) IsBought() bool {
	return b.bondIsBought.BondIsBought()
}

// CreatorUserID returns the ID of the user who created the Bond
func (b Bond) CreatorUserID() string {
	return b.bondCreatorUserId.ID()
}

// CurrentOwnerID returns the ID of the current owner of the Bond
func (b Bond) CurrentOwnerID() string {
	return b.bondCurrentOwnerId.ID()
}
