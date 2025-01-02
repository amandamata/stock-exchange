package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(
	id string,
	sellingOrder *Order,
	buyingOrer *Order,
	shares int,
	price float64,
) *Transaction {

	total := float64(shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrer,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}
