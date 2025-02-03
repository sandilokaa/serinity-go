package material

type MaterialFormatter struct {
	ID           int    `json:"id"`
	MaterialName string `json:"material_name"`
}

func FormatMaterial(material Material) MaterialFormatter {
	materialFormatter := MaterialFormatter{}
	materialFormatter.ID = material.ID
	materialFormatter.MaterialName = material.MaterialName

	return materialFormatter
}

func FormatMaterials(materials []Material) []MaterialFormatter {
	materialsFormatter := []MaterialFormatter{}

	for _, material := range materials {
		materialFormatter := FormatMaterial(material)
		materialsFormatter = append(materialsFormatter, materialFormatter)
	}

	return materialsFormatter
}

type MaterialDetailFormatter struct {
	ID           int    `json:"id"`
	MaterialName string `json:"material_name"`
}

type MaterialUserFormatter struct {
	Name string `json:"name"`
}

func (m *Material) FormatMaterialDetail(material Material) MaterialDetailFormatter {
	materialDetailFormatter := MaterialDetailFormatter{}
	materialDetailFormatter.ID = material.ID
	materialDetailFormatter.MaterialName = material.MaterialName

	return materialDetailFormatter
}

func UpdatedFormatMaterial(updatedMaterial Material, oldMaterial Material) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldMaterial.UserID != updatedMaterial.UserID {
		updatedFields["user_id"] = updatedMaterial.UserID
	}
	if oldMaterial.MaterialName != updatedMaterial.MaterialName {
		updatedFields["material_name"] = updatedMaterial.MaterialName
	}

	return updatedFields
}
