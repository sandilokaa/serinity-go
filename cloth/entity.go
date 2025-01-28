package cloth

import (
	"serinitystore/category"
	"serinitystore/material"
	sizechart "serinitystore/size-chart"
	"serinitystore/supplier"
	"serinitystore/user"
	"time"
)

type Cloth struct {
	ID          int
	UserID      int
	SupplierID  int
	MaterialID  int
	CategoryID  int
	SizeChartID int
	Name        string
	Price       string
	Description string
	Sale        bool
	NewArrival  bool
	ClothImages []ClothImage
	User        user.User
	Material    material.Material
	Supplier    supplier.Supplier
	Category    category.Category
	SizeChart   sizechart.SizeChart
	Variation   []ClothVariation
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ClothVariation struct {
	ID        int
	UserID    int
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
	UserID    int
	ClothID   int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User
}

func (ClothImage) TableName() string {
	return "ClothImages"
}
