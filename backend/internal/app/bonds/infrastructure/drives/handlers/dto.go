package handlers

import "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"

type (
	// BondRequest represnts a DTO(Data Transfer Object)
	BondRequest struct {
		ID             string  `json:"id"`
		Name           string  `json:"name"`
		QuantitySale   int64   `json:"quantity_sale"`
		SalesPrice     float64 `json:"sales_price"`
		IsBought       bool    `json:"is_bought"`
		CreatorUserId  string  `json:"creator_user_id"`
		CurrentOwnerId string  `json:"current_owner_id"`
	}

	// BondsRequest is a BondRequest type collection
	BondsRequest []*BondRequest
)

func (b BondRequest) ToBusiness() (*domain.Bond, error) {
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

func toRequest(bonds domain.Bonds) BondsRequest {
	bondsRequest := make(BondsRequest, 0, len(bonds))

	for _, bond := range bonds {
		bondsRequest = append(bondsRequest, &BondRequest{
			ID:             bond.ID(),
			Name:           bond.Name(),
			QuantitySale:   bond.QuantitySale(),
			SalesPrice:     bond.SalesPrice(),
			IsBought:       bond.IsBought(),
			CreatorUserId:  bond.CreatorUserID(),
			CurrentOwnerId: bond.CurrentOwnerID(),
		})
	}

	return bondsRequest
}
