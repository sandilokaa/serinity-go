package material

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAllMaterial(search string) ([]Material, error)
	FindMaterialById(ID int) (Material, error)
	SaveMaterial(material Material) (Material, error)
	UpdateMaterial(material Material) (Material, error)
	DeleteMaterial(ID int) (Material, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllMaterial(search string) ([]Material, error) {
	var materials []Material

	query := r.db
	if search != "" {
		query = query.Where("material_name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&materials).Error
	if err != nil {
		return materials, err
	}

	return materials, nil
}

func (r *repository) FindMaterialById(ID int) (Material, error) {
	var material Material

	err := r.db.Preload("User").Where("id = ?", ID).Find(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) SaveMaterial(material Material) (Material, error) {
	err := r.db.Create(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) UpdateMaterial(material Material) (Material, error) {
	err := r.db.Save(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}

func (r *repository) DeleteMaterial(ID int) (Material, error) {
	var material Material
	err := r.db.Where("id = ?", ID).Delete(&material).Error
	if err != nil {
		return material, err
	}

	return material, nil
}
