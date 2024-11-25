package category

import "gorm.io/gorm"

type Repository interface {
	SaveCategory(category Category) (Category, error)
	FindAllCategory(search string) ([]Category, error)
	FindCategoryByID(ID int) (Category, error)
	UpdateCategoryByID(category Category) (Category, error)
	DeleteCategoryByID(ID int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveCategory(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindAllCategory(search string) ([]Category, error) {
	var categories []Category

	query := r.db
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *repository) FindCategoryByID(ID int) (Category, error) {
	var category Category

	err := r.db.Preload("User").Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) UpdateCategoryByID(category Category) (Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) DeleteCategoryByID(ID int) (Category, error) {
	var category Category
	err := r.db.Where("id = ?", ID).Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
