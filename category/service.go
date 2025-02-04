package category

import (
	"fmt"
	"serinitystore/helper"
	"serinitystore/redis"
	"time"
)

type Service interface {
	CreateCategory(input CreateCategoryInput) (Category, error)
	FindAllCategory(search string) ([]Category, error)
	FindCategoryByID(input GetCategoryDetailInput) (Category, error)
	UpdateCategoryByID(inputID GetCategoryDetailInput, inputData UpdateCategoryInput) (Category, error)
	DeleteCategoryByID(inputID GetCategoryDetailInput) (Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCategory(input CreateCategoryInput) (Category, error) {
	category := Category{}
	category.UserID = input.User.ID
	category.Category = input.Category

	newCategory, err := s.repository.SaveCategory(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) FindAllCategory(search string) ([]Category, error) {
	redisClient := redis.GetRedisClient()
	var cacheKey string

	if search == "" {
		cacheKey = "categories:all"
	} else {
		cacheKey = fmt.Sprintf("categories:%s", search)
	}

	return helper.GetOrSetCache(redisClient, cacheKey, 5*time.Minute, func() ([]Category, error) {
		return s.repository.FindAllCategory(search)
	})
}

func (s *service) FindCategoryByID(input GetCategoryDetailInput) (Category, error) {
	redisClient := redis.GetRedisClient()
	cacheKey := fmt.Sprintf("category:%d", input.ID)

	return helper.GetOrSetCache(redisClient, cacheKey, 5*time.Minute, func() (Category, error) {
		return s.repository.FindCategoryByID(input.ID)
	})
}

func (s *service) UpdateCategoryByID(inputID GetCategoryDetailInput, inputData UpdateCategoryInput) (Category, error) {
	category, err := s.repository.FindCategoryByID(inputID.ID)
	if err != nil {
		return category, err
	}

	if inputData.Category != "" {
		category.Category = inputData.Category
	}

	updatedCategory, err := s.repository.UpdateCategoryByID(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *service) DeleteCategoryByID(inputID GetCategoryDetailInput) (Category, error) {
	category, err := s.repository.FindCategoryByID(inputID.ID)
	if err != nil {
		return category, err
	}

	deletedCategory, err := s.repository.DeleteCategoryByID(category.ID)
	if err != nil {
		return deletedCategory, err
	}

	return deletedCategory, nil
}
