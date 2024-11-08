package cloth

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SaveCloth(cloth Cloth) (Cloth, error)
	FindAllCloth(search string) ([]Cloth, error)
	FindClothByID(ID int) (Cloth, error)
	UpdateClothByID(cloth Cloth) (Cloth, error)
	UpdateStockByClothID(clothID int, newStock int) error
	DeleteClothById(ID int) (Cloth, error)
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

func (r *repository) FindAllCloth(search string) ([]Cloth, error) {
	var cloths []Cloth

	query := r.db
	if search != "" {
		query = query.Preload("ClothImages", "cloth_images.is_primary = 1").Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Preload("ClothImages", "cloth_images.is_primary = 1").Find(&cloths).Error
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

func (r *repository) UpdateClothByID(cloth Cloth) (Cloth, error) {
	err := r.db.Save(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) UpdateStockByClothID(clothID int, newStock int) error {
	var cloth Cloth
	if err := r.db.First(&cloth, clothID).Error; err != nil {
		return err
	}

	cloth.Stock = newStock
	return r.db.Save(&cloth).Error
}

func (r *repository) DeleteClothById(ID int) (Cloth, error) {
	var cloth Cloth
	err := r.db.Where("id ? =", ID).Delete(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
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
