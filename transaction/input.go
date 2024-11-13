package transaction

import (
	"cheggstore/user"
)

type CreateTransactionInput struct {
	User             user.User
	ClothID          int `json:"cloth_id" binding:"required"`
	ClothVariationID int `json:"cloth_variation_id" binding:"required"`
	Quantity         int `json:"quantity" binding:"required"`
}

type TransactionInputDetail struct {
	ID int `uri:"id" binding:"required"`
}
