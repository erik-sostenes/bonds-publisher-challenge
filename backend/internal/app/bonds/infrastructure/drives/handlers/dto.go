package handlers

import "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"

// BondRequest represnts a DTO(Data Transfer Object)
type BondRequest struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	QuantitySale   int64   `json:"quantity_sale"`
	SalesPrice     float64 `json:"sales_price"`
	IsBought       bool    `json:"is_bought"`
	CreatorUserId  string  `json:"creator_user_id"`
	CurrentOwnerId string  `json:"current_owner_id"`
}

func (b BondRequest) toBusiness() (*domain.Bond, error) {
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
