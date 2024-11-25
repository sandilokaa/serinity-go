package category

import "cheggstore/user"

type CreateCategoryInput struct {
	User     user.User
	Category string `json:"category"`
}

type UpdateCategoryInput struct {
	User     user.User
	Category string `json:"category"`
}

type GetCategoryDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
