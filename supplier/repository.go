package supplier

import "gorm.io/gorm"

type Repository interface {
	SaveSupplier(supplier Supplier) (Supplier, error)
	FindAllSupplier(search string) ([]Supplier, error)
	FindSupplierByID(ID int) (Supplier, error)
	UpdateSupplierByID(supplier Supplier) (Supplier, error)
	DeleteSupplierByID(ID int) (Supplier, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveSupplier(supplier Supplier) (Supplier, error) {
	err := r.db.Create(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FindAllSupplier(search string) ([]Supplier, error) {
	var suppliers []Supplier

	query := r.db
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Find(&suppliers).Error
	if err != nil {
		return suppliers, err
	}

	return suppliers, nil
}

func (r *repository) FindSupplierByID(ID int) (Supplier, error) {
	var supplier Supplier

	err := r.db.Preload("User").Where("id = ?", ID).Find(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) UpdateSupplierByID(supplier Supplier) (Supplier, error) {
	err := r.db.Save(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) DeleteSupplierByID(ID int) (Supplier, error) {
	var supplier Supplier
	err := r.db.Where("id = ?", ID).Delete(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}
