package supplier

import "cheggstore/user"

type CreateSupplierInput struct {
	User    user.User
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Postal  string `json:"postal" binding:"required"`
}

type UpdateSupplierInput struct {
	User    user.User
	Name    string `json:"name"`
	Address string `json:"address"`
	Postal  string `json:"postal"`
}

type GetSupplierDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
