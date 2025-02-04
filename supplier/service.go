package supplier

import (
	"fmt"
	"serinitystore/helper"
	"serinitystore/redis"
	"time"
)

type Service interface {
	CreateSupplier(input CreateSupplierInput) (Supplier, error)
	FindAllSupplier(search string) ([]Supplier, error)
	FindSupplierByID(input GetSupplierDetailInput) (Supplier, error)
	UpdateSupplierByID(inputID GetSupplierDetailInput, inputData UpdateSupplierInput) (Supplier, error)
	DeleteSupplierByID(inputID GetSupplierDetailInput) (Supplier, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSupplier(input CreateSupplierInput) (Supplier, error) {
	supplier := Supplier{}
	supplier.UserID = input.User.ID
	supplier.Name = input.Name
	supplier.Address = input.Address
	supplier.Postal = input.Postal

	newSupplier, err := s.repository.SaveSupplier(supplier)
	if err != nil {
		return newSupplier, err
	}

	return newSupplier, nil
}

func (s *service) FindAllSupplier(search string) ([]Supplier, error) {
	redisClient := redis.GetRedisClient()
	var cacheKey string

	if search == "" {
		cacheKey = "suppliers:all"
	} else {
		cacheKey = fmt.Sprintf("suppliers:%s", search)
	}

	return helper.GetOrSetCache(redisClient, cacheKey, 5*time.Minute, func() ([]Supplier, error) {
		return s.repository.FindAllSupplier(search)
	})
}

func (s *service) FindSupplierByID(input GetSupplierDetailInput) (Supplier, error) {
	redisClient := redis.GetRedisClient()
	cacheKey := fmt.Sprintf("supplier:%d", input.ID)

	return helper.GetOrSetCache(redisClient, cacheKey, 5*time.Minute, func() (Supplier, error) {
		return s.repository.FindSupplierByID(input.ID)
	})
}

func (s *service) UpdateSupplierByID(inputID GetSupplierDetailInput, inputData UpdateSupplierInput) (Supplier, error) {
	supplier, err := s.repository.FindSupplierByID(inputID.ID)
	if err != nil {
		return supplier, err
	}

	if inputData.Name != "" {
		supplier.Name = inputData.Name
	}
	if inputData.Address != "" {
		supplier.Address = inputData.Address
	}
	if inputData.Postal != "" {
		supplier.Postal = inputData.Postal
	}

	updatedSupplier, err := s.repository.UpdateSupplierByID(supplier)
	if err != nil {
		return updatedSupplier, err
	}

	return updatedSupplier, nil
}

func (s *service) DeleteSupplierByID(inputID GetSupplierDetailInput) (Supplier, error) {
	supplier, err := s.repository.FindSupplierByID(inputID.ID)
	if err != nil {
		return supplier, err
	}

	deletedSupplier, err := s.repository.DeleteSupplierByID(supplier.ID)
	if err != nil {
		return deletedSupplier, err
	}

	return deletedSupplier, nil
}
