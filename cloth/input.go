package cloth

import (
	"serinitystore/user"
)

type CreateClothInput struct {
	User        user.User
	MaterialID  int                   `json:"material_id" binding:"required"`
	SupplierID  int                   `json:"supplier_id" binding:"required"`
	CategoryID  int                   `json:"category_id" binding:"required"`
	SizeChartID int                   `json:"size_chart_id" binding:"required"`
	Name        string                `json:"name" binding:"required"`
	Price       string                `json:"price" binding:"required"`
	Description string                `json:"description" binding:"required"`
	Sale        bool                  `json:"sale"`
	NewArrival  bool                  `json:"new_arrival"`
	Variations  []ClothVariationInput `json:"variations"`
}

type ClothVariationInput struct {
	User  user.User
	Size  string `json:"size" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type ClothInputDetail struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateClothInput struct {
	User        user.User
	MaterialID  int    `json:"material_id"`
	SupplierID  int    `json:"supplier_id"`
	CategoryID  int    `json:"category_id"`
	SizeChartID int    `json:"size_chart_id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Sale        bool   `json:"sale"`
	NewArrival  bool   `json:"new_arrival"`
	Size        string `json:"size"`
	Stock       int    `json:"stock"`
}

type UpdateClothVariationInput struct {
	Size  string `json:"size"`
	Color string `json:"color"`
	Stock int    `json:"stock"`
}

type CreateClothImageInput struct {
	ClothID   int  `form:"cloth_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}
