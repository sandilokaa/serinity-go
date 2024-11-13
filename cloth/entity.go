package cloth

import (
	"cheggstore/material"
	"cheggstore/supplier"
	"cheggstore/user"
	"time"
)

type Cloth struct {
	ID          int
	UserID      int
	SupplierID  int
	MaterialID  int
	Name        string
	Price       string
	Description string
	ClothImages []ClothImage
	User        user.User
	Material    material.Material
	Supplier    supplier.Supplier
	Variation   []ClothVariation
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ClothVariation struct {
	ID        int
	ClothID   int
	Size      string
	Stock     int
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ClothVariation) TableName() string {
	return "ClothVariations"
}

type ClothImage struct {
	ID        int
	ClothID   int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ClothImage) TableName() string {
	return "ClothImages"
}
