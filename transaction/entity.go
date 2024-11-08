package transaction

import (
	"cheggstore/cloth"
	"cheggstore/user"
	"time"
)

type Transaction struct {
	ID         int
	UserID     int
	ClothID    int
	Quantity   int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Cloth      cloth.Cloth
}
