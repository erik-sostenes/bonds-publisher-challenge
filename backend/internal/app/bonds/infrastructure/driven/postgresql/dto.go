package postgresql

import "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"

type (
	// BondSchema represnts a DTO(Data Transfer Object)
	BondSchema struct {
		ID             string
		Name           string
		QuantitySale   int64
		SalesPrice     float64
		IsBought       bool
		CreatorUserId  string
		CurrentOwnerId string
	}

	// BondsSchema is a BondSchema type collection
	BondsSchema []*BondSchema
)

func (b *BondsSchema) ToBusiness() (domain.Bonds, error) {
	bonds := make(domain.Bonds, 0, len(*b))

	for _, bondSchema := range *b {
		bond, err := bondSchema.ToBusiness()
		if err != nil {
			return nil, err
		}

		bonds = append(bonds, bond)
	}
	return bonds, nil
}

func (b *BondSchema) ToBusiness() (*domain.Bond, error) {
	return domain.NewBond(
		b.ID,
		b.Name,
		b.CreatorUserId,
		b.CurrentOwnerId,
		b.IsBought,
		b.QuantitySale,
		b.SalesPrice,
	)
}
