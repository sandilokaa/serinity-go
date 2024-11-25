package category

type CategoryFormatter struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Category string `json:"category"`
}

func FormatCategory(category Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.ID = category.ID
	categoryFormatter.UserID = category.UserID
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
	ID       int                   `json:"id"`
	UserID   int                   `json:"user_id"`
	Category string                `json:"category"`
	User     CategoryUserFormatter `json:"user"`
}

type CategoryUserFormatter struct {
	Name string `json:"name"`
}

func (s *Category) FormatCategoryDetail(category Category) CategoryDetailFormatter {

	categoryDetailFormatter := CategoryDetailFormatter{}
	categoryDetailFormatter.ID = category.ID
	categoryDetailFormatter.Category = category.Category

	user := category.User
	categoryUserFormatter := CategoryUserFormatter{}
	categoryUserFormatter.Name = user.Name

	categoryDetailFormatter.User = categoryUserFormatter

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
