package material

import "serinitystore/user"

type GetMaterialDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMaterialInput struct {
	User         user.User
	MaterialName string `json:"material_name" binding:"required"`
}
