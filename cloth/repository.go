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
	DeleteClothById(ID int) (Cloth, error)
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
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&cloths).Error
	if err != nil {
		return cloths, err
	}

	return cloths, nil
}

func (r *repository) FindClothByID(ID int) (Cloth, error) {
	var cloth Cloth

	err := r.db.Preload(clause.Associations).Where("id = ?", ID).Find(&cloth).Error
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

func (r *repository) DeleteClothById(ID int) (Cloth, error) {
	var cloth Cloth
	err := r.db.Where("id ? =", ID).Delete(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}
