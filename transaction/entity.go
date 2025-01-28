package transaction

import (
	"serinitystore/cloth"
	"serinitystore/user"
	"time"
)

type Transaction struct {
	ID               int
	UserID           int
	ClothID          int
	ClothVariationID int
	Quantity         int
	Amount           int
	Status           string
	Code             string
	PaymentURL       string
	User             user.User
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Cloth            cloth.Cloth
}
