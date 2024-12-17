package cloth

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SaveCloth(cloth Cloth) (Cloth, error)
	SaveClothVariation(clothVariation ClothVariation) (ClothVariation, error)
	FindAllCloth(name string, category string) ([]Cloth, error)
	FindClothByID(ID int) (Cloth, error)
	FindClothVariationByID(ID int) (ClothVariation, error)
	UpdateClothByID(cloth Cloth) (Cloth, error)
	UpdateClothVariationByID(clothVariation ClothVariation) (ClothVariation, error)
	UpdateStockByClothID(ID int, newStock int) error
	DeleteClothById(ID int) (Cloth, error)
	DeleteClothVariationByClothId(clothID int) (ClothVariation, error)
	CreateClothImage(clothImage ClothImage) (ClothImage, error)
	MarkAllImagesAsNonPrimary(clothID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveCloth(cloth Cloth) (Cloth, error) {
	err := r.db.Create(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) SaveClothVariation(clothVariation ClothVariation) (ClothVariation, error) {
	err := r.db.Create(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) FindAllCloth(name string, category string) ([]Cloth, error) {
	var cloths []Cloth

	query := r.db.Preload("ClothImages", "ClothImages.is_primary = 1").Preload("Category")

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if category != "" {
		query = query.Joins("JOIN categories ON categories.id = cloths.category_id").
			Where("categories.category LIKE ?", "%"+category+"%")
	}

	err := query.Find(&cloths).Error
	if err != nil {
		return cloths, err
	}

	return cloths, nil
}

func (r *repository) FindClothByID(ID int) (Cloth, error) {
	var cloth Cloth

	err := r.db.Preload(clause.Associations).Preload("ClothImages").Where("id = ?", ID).Find(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) FindClothVariationByID(ID int) (ClothVariation, error) {
	var clothVariation ClothVariation

	err := r.db.Where("id = ?", ID).Find(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) UpdateClothByID(cloth Cloth) (Cloth, error) {
	err := r.db.Save(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) UpdateClothVariationByID(clothVariation ClothVariation) (ClothVariation, error) {
	err := r.db.Save(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) UpdateStockByClothID(ID int, newStock int) error {
	var clothVariation ClothVariation
	if err := r.db.First(&clothVariation, ID).Error; err != nil {
		return err
	}

	clothVariation.Stock = newStock
	return r.db.Save(&clothVariation).Error
}

func (r *repository) DeleteClothById(ID int) (Cloth, error) {
	var cloth Cloth
	err := r.db.Where("id = ?", ID).Delete(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) DeleteClothVariationByClothId(clothID int) (ClothVariation, error) {
	var clothVariation ClothVariation
	err := r.db.Where("cloth_id = ?", clothID).Delete(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) CreateClothImage(clothImage ClothImage) (ClothImage, error) {
	err := r.db.Create(&clothImage).Error
	if err != nil {
		return clothImage, err
	}

	return clothImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(clothID int) (bool, error) {
	err := r.db.Model(&ClothImage{}).Where("cloth_id = ?", clothID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
