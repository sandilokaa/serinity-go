package cloth

type ClothFormatter struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	MaterialID  int    `json:"material_id"`
	SupplierID  int    `json:"supplier_id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Stock       int    `json:"stock"`
	Color       string `json:"color"`
	ImageURL    string `json:"image_url"`
}

func FormatCloth(cloth Cloth) ClothFormatter {
	clothFormatter := ClothFormatter{}
	clothFormatter.ID = cloth.ID
	clothFormatter.UserID = cloth.UserID
	clothFormatter.MaterialID = cloth.MaterialID
	clothFormatter.SupplierID = cloth.SupplierID
	clothFormatter.Name = cloth.Name
	clothFormatter.Price = cloth.Price
	clothFormatter.Description = cloth.Description
	clothFormatter.Size = cloth.Size
	clothFormatter.Stock = cloth.Stock
	clothFormatter.Color = cloth.Color
	clothFormatter.ImageURL = ""

	if len(cloth.ClothImages) > 0 {
		clothFormatter.ImageURL = cloth.ClothImages[0].FileName
	}

	return clothFormatter
}

func FormatCloths(cloths []Cloth) []ClothFormatter {
	clothsFormatter := []ClothFormatter{}

	for _, supplier := range cloths {
		clothFormatter := FormatCloth(supplier)
		clothsFormatter = append(clothsFormatter, clothFormatter)
	}

	return clothsFormatter
}

type ClothDetailFormatter struct {
	ID          int                    `json:"id"`
	UserID      int                    `json:"user_id"`
	MaterialID  int                    `json:"material_id"`
	SupplierID  int                    `json:"supplier_id"`
	Name        string                 `json:"name"`
	Price       string                 `json:"price"`
	Description string                 `json:"description"`
	Size        string                 `json:"size"`
	Color       string                 `json:"color"`
	Stock       int                    `json:"stock"`
	User        ClothUserFormatter     `json:"user"`
	Material    ClothMaterialFormatter `json:"material"`
	Supplier    ClothSupplierFormatter `json:"supplier"`
	Images      []ClothImageFormatter  `json:"images"`
}

type ClothUserFormatter struct {
	Name string `json:"name"`
}

type ClothMaterialFormatter struct {
	MaterialName string `json:"material_name"`
}

type ClothSupplierFormatter struct {
	Name string `json:"name"`
}

type ClothImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func (c *Cloth) FormatClothDetail(cloth Cloth) ClothDetailFormatter {
	clothDetailFormatter := ClothDetailFormatter{}
	clothDetailFormatter.ID = cloth.ID
	clothDetailFormatter.UserID = cloth.UserID
	clothDetailFormatter.MaterialID = cloth.MaterialID
	clothDetailFormatter.SupplierID = cloth.SupplierID
	clothDetailFormatter.Name = cloth.Name
	clothDetailFormatter.Price = cloth.Price
	clothDetailFormatter.Color = cloth.Color
	clothDetailFormatter.Stock = cloth.Stock
	clothDetailFormatter.Size = cloth.Size
	clothDetailFormatter.Description = cloth.Description

	user := cloth.User
	clothUserFormatter := ClothUserFormatter{}
	clothUserFormatter.Name = user.Name

	material := cloth.Material
	clothMaterialFormatter := ClothMaterialFormatter{}
	clothMaterialFormatter.MaterialName = material.MaterialName

	supplier := cloth.Supplier
	clothSupplierFormatter := ClothSupplierFormatter{}
	clothSupplierFormatter.Name = supplier.Name

	clothDetailFormatter.User = clothUserFormatter
	clothDetailFormatter.Material = clothMaterialFormatter
	clothDetailFormatter.Supplier = clothSupplierFormatter

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
	if oldCloth.Name != updatedCloth.Name {
		updatedFields["name"] = updatedCloth.Name
	}
	if oldCloth.Price != updatedCloth.Price {
		updatedFields["price"] = updatedCloth.Price
	}
	if oldCloth.Description != updatedCloth.Description {
		updatedFields["description"] = updatedCloth.Description
	}
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

func UpdateFormatClothImage(updatedClothImage ClothImage, oldClothImage ClothImage) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldClothImage.ClothID != updatedClothImage.ClothID {
		updatedFields["cloth_id"] = updatedClothImage.ClothID
	}

	if oldClothImage.IsPrimary != updatedClothImage.IsPrimary {
		updatedFields["is_primary"] = updatedClothImage.IsPrimary
	}

	if oldClothImage.FileName != updatedClothImage.FileName {
		updatedFields["file"] = updatedClothImage.FileName
	}

	return updatedFields
}
