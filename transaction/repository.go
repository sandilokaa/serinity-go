package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAllTransaction(search string) ([]Transaction, error)
	GetTransactionByUserID(requestedUserID int) ([]Transaction, error)
	GetTransactionByID(ID int) (Transaction, error)
	GetTransactionUserIDByID(ID int, userID int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllTransaction(search string) ([]Transaction, error) {
	var transactions []Transaction
	query := r.db
	if search != "" {
		query = query.Preload("Cloth.ClothImages", "cloth_images.is_primary = 1").Where("created_at LIKE ?", "%"+search+"%")
	}

	err := query.Preload("Cloth.ClothImages", "cloth_images.is_primary = 1").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetTransactionByUserID(requestedUserID int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("Cloth.ClothImages", "cloth_images.is_primary = 1").Where("user_id = ?", requestedUserID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetTransactionByID(ID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Preload("User").Preload("Cloth.Material").Preload("Cloth.ClothImages", "cloth_images.is_primary = 1").Where("id = ?", ID).First(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetTransactionUserIDByID(ID int, userID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Preload("User").Preload("Cloth.Material").Preload("Cloth.ClothImages", "cloth_images.is_primary = 1").Where("id = ? AND user_id = ?", ID, userID).First(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
