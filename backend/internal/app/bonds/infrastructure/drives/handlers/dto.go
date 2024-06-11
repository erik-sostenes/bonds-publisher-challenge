package handlers

type BondRequest struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	QuantitySale   int64   `json:"quantity_sale"`
	SalesPrice     float64 `json:"sales_price"`
	IsBought       bool    `json:"is_bought"`
	CreatorUserId  string  `json:"creator_user_id"`
	CurrentOwnerId string  `json:"current_owner_id"`
}
