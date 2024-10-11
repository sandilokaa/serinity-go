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
	Color       string
	Price       string
	Size        string
	Description string
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ClothImages []ClothImage
	User        user.User
	Material    material.Material
	Supplier    supplier.Supplier
}

type ClothImage struct {
	ID        int
	ClothID   int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
