package sizechart

import (
	"os"

	"gorm.io/gorm"
)

type Repository interface {
	SaveSizeChart(sizeChart SizeChart) (SizeChart, error)
	UpdateSizeChartByID(sizeChart SizeChart) (SizeChart, error)
	FindSizeChartByID(ID int) (SizeChart, error)
	DeleteImage(fileLocation string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveSizeChart(sizeChart SizeChart) (SizeChart, error) {
	err := r.db.Create(&sizeChart).Error
	if err != nil {
		return sizeChart, err
	}

	return sizeChart, nil
}

func (r *repository) FindSizeChartByID(ID int) (SizeChart, error) {
	var sizeChart SizeChart
	err := r.db.Where("id = ?", ID).Find(&sizeChart).Error
	if err != nil {
		return sizeChart, err
	}

	return sizeChart, nil
}

func (r *repository) UpdateSizeChartByID(sizeChart SizeChart) (SizeChart, error) {
	err := r.db.Save(&sizeChart).Error
	if err != nil {
		return sizeChart, err
	}

	return sizeChart, nil
}

func (r *repository) DeleteImage(fileLocation string) error {
	err := os.Remove(fileLocation)
	if err != nil {
		return err
	}

	return nil
}
