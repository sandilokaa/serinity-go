package cloth

type ClothFormatter struct {
	ID         int                    `json:"id"`
	Name       string                 `json:"name"`
	Price      string                 `json:"price"`
	Sale       bool                   `json:"sale"`
	NewArrival bool                   `json:"new_arrival"`
	ImageURL   string                 `json:"image_url"`
	Category   ClothCategoryFormatter `json:"category"`
}

type ClothCategoryFormatter struct {
	Category string `json:"category"`
}

func FormatCloth(cloth Cloth) ClothFormatter {
	clothFormatter := ClothFormatter{}
	clothFormatter.ID = cloth.ID
	clothFormatter.Name = cloth.Name
	clothFormatter.Price = cloth.Price
	clothFormatter.Sale = cloth.Sale
	clothFormatter.NewArrival = cloth.NewArrival
	clothFormatter.ImageURL = ""

	category := cloth.Category
	clothCategoryFormatter := ClothCategoryFormatter{}
	clothCategoryFormatter.Category = category.Category

	clothFormatter.Category = clothCategoryFormatter

	if len(cloth.ClothImages) > 0 {
		clothFormatter.ImageURL = cloth.ClothImages[0].FileName
	}

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
	ID          int                             `json:"id"`
	Name        string                          `json:"name"`
	Price       string                          `json:"price"`
	Description string                          `json:"description"`
	Sale        bool                            `json:"sale"`
	NewArrival  bool                            `json:"new_arrival"`
	Material    ClothDetailMaterialFormatter    `json:"material"`
	SizeChart   ClothDetailSizeChartFormatter   `json:"size_chart"`
	Category    ClothDetailCategoryFormatter    `json:"category"`
	Images      []ClothDetailImageFormatter     `json:"images"`
	Variations  []ClothDetailVariationFormatter `json:"variations"`
}

type ClothDetailMaterialFormatter struct {
	MaterialName string `json:"material_name"`
}

type ClothDetailCategoryFormatter struct {
	Category string `json:"category"`
}

type ClothDetailImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type ClothDetailSizeChartFormatter struct {
	Name     string `json:"name"`
	FileName string `json:"file_name"`
}

type ClothDetailVariationFormatter struct {
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
	clothMaterialFormatter := ClothDetailMaterialFormatter{}
	clothMaterialFormatter.MaterialName = material.MaterialName

	clothDetailFormatter.Material = clothMaterialFormatter

	category := cloth.Category
	clothCategoryFormatter := ClothDetailCategoryFormatter{}
	clothCategoryFormatter.Category = category.Category

	clothDetailFormatter.Category = clothCategoryFormatter

	sizechart := cloth.SizeChart
	sizeChartFormatter := ClothDetailSizeChartFormatter{}
	sizeChartFormatter.Name = sizechart.Name
	sizeChartFormatter.FileName = sizechart.FileName

	clothDetailFormatter.SizeChart = sizeChartFormatter

	for _, variation := range cloth.Variation {
		clothVariationFormatter := ClothDetailVariationFormatter{
			Color: variation.Color,
			Size:  variation.Size,
			Stock: variation.Stock,
		}

		clothDetailFormatter.Variations = append(clothDetailFormatter.Variations, clothVariationFormatter)
	}

	images := []ClothDetailImageFormatter{}

	for _, image := range cloth.ClothImages {
		clothImageFormatter := ClothDetailImageFormatter{}
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
	if oldCloth.SizeChartID != updatedCloth.SizeChartID {
		updatedFields["size_chart_id"] = updatedCloth.SizeChartID
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
