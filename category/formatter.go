package category

type CategoryFormatter struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

func FormatCategory(category Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.ID = category.ID
	categoryFormatter.Category = category.Category

	return categoryFormatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	categoriesFormatter := []CategoryFormatter{}

	for _, category := range categories {
		categoryFormatter := FormatCategory(category)
		categoriesFormatter = append(categoriesFormatter, categoryFormatter)
	}

	return categoriesFormatter
}

type CategoryDetailFormatter struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

type CategoryUserFormatter struct {
	Name string `json:"name"`
}

func (s *Category) FormatCategoryDetail(category Category) CategoryDetailFormatter {

	categoryDetailFormatter := CategoryDetailFormatter{}
	categoryDetailFormatter.ID = category.ID
	categoryDetailFormatter.Category = category.Category

	return categoryDetailFormatter
}

func UpdatedFormatCategory(updatedCategory Category, oldCategory Category) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldCategory.ID != updatedCategory.ID {
		updatedFields["id"] = updatedCategory.ID
	}
	if oldCategory.UserID != updatedCategory.UserID {
		updatedFields["user_id"] = updatedCategory.UserID
	}
	if oldCategory.Category != updatedCategory.Category {
		updatedFields["category"] = updatedCategory.Category
	}

	return updatedFields
}
