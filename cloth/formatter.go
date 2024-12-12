package cloth

type ClothFormatter struct {
	ID          int         `json:"id"`
	UserID      int         `json:"user_id"`
	MaterialID  int         `json:"material_id"`
	SupplierID  int         `json:"supplier_id"`
	CategoryID  int         `json:"category_id"`
	Name        string      `json:"name"`
	Price       string      `json:"price"`
	Description string      `json:"description"`
	Sale        bool        `json:"sale"`
	NewArrival  bool        `json:"new_arrival"`
	ImageURL    string      `json:"image_url"`
	Variations  interface{} `json:"variations"`
}

func FormatCloth(cloth Cloth) ClothFormatter {
	clothFormatter := ClothFormatter{}
	clothFormatter.ID = cloth.ID
	clothFormatter.UserID = cloth.UserID
	clothFormatter.MaterialID = cloth.MaterialID
	clothFormatter.SupplierID = cloth.SupplierID
	clothFormatter.CategoryID = cloth.CategoryID
	clothFormatter.Name = cloth.Name
	clothFormatter.Price = cloth.Price
	clothFormatter.Description = cloth.Description
	clothFormatter.Sale = cloth.Sale
	clothFormatter.NewArrival = cloth.NewArrival
	clothFormatter.ImageURL = ""

	if len(cloth.ClothImages) > 0 {
		clothFormatter.ImageURL = cloth.ClothImages[0].FileName
	}

	clothFormatter.Variations = "Success"

	return clothFormatter
}

func FormatCloths(cloths []Cloth) []ClothFormatter {
	clothsFormatter := []ClothFormatter{}

	for _, cloth := range cloths {
		clothFormatter := FormatCloth(cloth)
		clothsFormatter = append(clothsFormatter, clothFormatter)
	}

	return clothsFormatter
}

type ClothDetailFormatter struct {
	ID          int                       `json:"id"`
	Name        string                    `json:"name"`
	Price       string                    `json:"price"`
	Description string                    `json:"description"`
	Sale        bool                      `json:"sale"`
	NewArrival  bool                      `json:"new_arrival"`
	Material    ClothMaterialFormatter    `json:"material"`
	Category    ClothCategoryFormatter    `json:"category"`
	Images      []ClothImageFormatter     `json:"images"`
	Variations  []ClothVariationFormatter `json:"variations"`
}

type ClothMaterialFormatter struct {
	MaterialName string `json:"material_name"`
}

type ClothCategoryFormatter struct {
	Category string `json:"category"`
}

type ClothImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type ClothVariationFormatter struct {
	Size  string `json:"size"`
	Stock int    `json:"stock"`
	Color string `json:"color"`
}

func (c *Cloth) FormatClothDetail(cloth Cloth) ClothDetailFormatter {
	clothDetailFormatter := ClothDetailFormatter{}
	clothDetailFormatter.ID = cloth.ID
	clothDetailFormatter.Name = cloth.Name
	clothDetailFormatter.Price = cloth.Price
	clothDetailFormatter.Description = cloth.Description
	clothDetailFormatter.Sale = cloth.Sale
	clothDetailFormatter.NewArrival = cloth.NewArrival

	material := cloth.Material
	clothMaterialFormatter := ClothMaterialFormatter{}
	clothMaterialFormatter.MaterialName = material.MaterialName

	clothDetailFormatter.Material = clothMaterialFormatter

	category := cloth.Category
	clothCategoryFormatter := ClothCategoryFormatter{}
	clothCategoryFormatter.Category = category.Category

	clothDetailFormatter.Category = clothCategoryFormatter

	for _, variation := range cloth.Variation {
		clothVariationFormatter := ClothVariationFormatter{
			Color: variation.Color,
			Size:  variation.Size,
			Stock: variation.Stock,
		}

		clothDetailFormatter.Variations = append(clothDetailFormatter.Variations, clothVariationFormatter)
	}

	images := []ClothImageFormatter{}

	for _, image := range cloth.ClothImages {
		clothImageFormatter := ClothImageFormatter{}
		clothImageFormatter.ImageUrl = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		clothImageFormatter.IsPrimary = isPrimary

		images = append(images, clothImageFormatter)
	}

	clothDetailFormatter.Images = images

	return clothDetailFormatter
}

func UpdatedFormatCloth(updatedCloth Cloth, oldCloth Cloth) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldCloth.UserID != updatedCloth.UserID {
		updatedFields["user_id"] = updatedCloth.UserID
	}
	if oldCloth.MaterialID != updatedCloth.MaterialID {
		updatedFields["material_id"] = updatedCloth.MaterialID
	}
	if oldCloth.SupplierID != updatedCloth.SupplierID {
		updatedFields["supplier_id"] = updatedCloth.SupplierID
	}
	if oldCloth.CategoryID != updatedCloth.CategoryID {
		updatedFields["category_id"] = updatedCloth.CategoryID
	}
	if oldCloth.Name != updatedCloth.Name {
		updatedFields["name"] = updatedCloth.Name
	}
	if oldCloth.Price != updatedCloth.Price {
		updatedFields["price"] = updatedCloth.Price
	}
	if oldCloth.Description != updatedCloth.Description {
		updatedFields["description"] = updatedCloth.Description
	}
	if oldCloth.Sale != updatedCloth.Sale {
		updatedFields["sale"] = updatedCloth.Sale
	}
	if oldCloth.NewArrival != updatedCloth.NewArrival {
		updatedFields["new_arrival"] = updatedCloth.NewArrival
	}

	return updatedFields
}

func UpdatedClothVariationFormatCloth(updatedCloth ClothVariation, oldCloth ClothVariation) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldCloth.Size != updatedCloth.Size {
		updatedFields["size"] = updatedCloth.Size
	}
	if oldCloth.Color != updatedCloth.Color {
		updatedFields["color"] = updatedCloth.Color
	}
	if oldCloth.Stock != updatedCloth.Stock {
		updatedFields["stock"] = updatedCloth.Stock
	}

	return updatedFields
}
