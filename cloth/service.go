package cloth

import "errors"

type Service interface {
	SaveCloth(input CreateClothInput) (Cloth, error)
	FindAllCloth(name string, category string) ([]Cloth, error)
	FindClothByID(input ClothInputDetail) (Cloth, error)
	FindClothVariationByID(input ClothInputDetail) (ClothVariation, error)
	UpdateClothByID(inputID ClothInputDetail, inputData UpdateClothInput) (Cloth, error)
	UpdateClothVariationByID(inputID ClothInputDetail, inputData UpdateClothVariationInput) (ClothVariation, error)
	DeleteClothByID(inputID ClothInputDetail) (Cloth, error)
	CreateClothImage(input CreateClothImageInput, fileLocation string) (ClothImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveCloth(input CreateClothInput) (Cloth, error) {
	cloth := Cloth{}
	cloth.Name = input.Name
	cloth.Price = input.Price
	cloth.Description = input.Description
	cloth.Sale = input.Sale
	cloth.NewArrival = input.NewArrival
	cloth.UserID = input.User.ID
	cloth.MaterialID = input.MaterialID
	cloth.SupplierID = input.SupplierID
	cloth.CategoryID = input.CategoryID

	newCloth, err := s.repository.SaveCloth(cloth)
	if err != nil {
		return newCloth, err
	}

	for _, variation := range input.Variations {
		clothVariation := ClothVariation{
			UserID:  newCloth.UserID,
			ClothID: newCloth.ID,
			Size:    variation.Size,
			Color:   variation.Color,
			Stock:   variation.Stock,
		}

		_, err := s.repository.SaveClothVariation(clothVariation)
		if err != nil {
			return newCloth, err
		}
	}

	return newCloth, nil
}

func (s *service) FindAllCloth(name string, category string) ([]Cloth, error) {
	cloths, err := s.repository.FindAllCloth(name, category)
	if err != nil {
		return cloths, err
	}

	return cloths, nil
}

func (s *service) FindClothByID(input ClothInputDetail) (Cloth, error) {
	cloth, err := s.repository.FindClothByID(input.ID)
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (s *service) FindClothVariationByID(input ClothInputDetail) (ClothVariation, error) {
	clothVariation, err := s.repository.FindClothVariationByID(input.ID)
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (s *service) UpdateClothByID(inputID ClothInputDetail, inputData UpdateClothInput) (Cloth, error) {

	cloth, err := s.repository.FindClothByID(inputID.ID)
	if err != nil {
		return cloth, err
	}

	if inputData.Name != "" {
		cloth.Name = inputData.Name
	}

	if inputData.Price != "" {
		cloth.Price = inputData.Price
	}

	if inputData.Description != "" {
		cloth.Description = inputData.Description
	}

	if inputData.Sale {
		cloth.Sale = inputData.Sale
	}

	if inputData.NewArrival {
		cloth.NewArrival = inputData.NewArrival
	}

	if inputData.MaterialID != 0 {
		cloth.MaterialID = inputData.MaterialID
	}

	if inputData.SupplierID != 0 {
		cloth.SupplierID = inputData.SupplierID
	}

	if inputData.CategoryID != 0 {
		cloth.CategoryID = inputData.CategoryID
	}

	updatedCloth, err := s.repository.UpdateClothByID(cloth)
	if err != nil {
		return updatedCloth, err
	}

	return updatedCloth, nil

}

func (s *service) UpdateClothVariationByID(inputID ClothInputDetail, inputData UpdateClothVariationInput) (ClothVariation, error) {
	clothVariation, err := s.repository.FindClothVariationByID(inputID.ID)
	if err != nil {
		return clothVariation, err
	}

	if inputData.Size != "" {
		clothVariation.Size = inputData.Size
	}

	if inputData.Stock != 0 {
		clothVariation.Stock = inputData.Stock
	}

	if inputData.Color != "" {
		clothVariation.Color = inputData.Color
	}

	updatedClothVariation, err := s.repository.UpdateClothVariationByID(clothVariation)
	if err != nil {
		return updatedClothVariation, err
	}

	return updatedClothVariation, nil
}

func (s *service) DeleteClothByID(inputID ClothInputDetail) (Cloth, error) {
	cloth, err := s.repository.FindClothByID(inputID.ID)
	if err != nil {
		return cloth, err
	}

	_, err = s.repository.DeleteClothVariationByClothId(cloth.ID)
	if err != nil {
		return cloth, err
	}

	deletedCloth, err := s.repository.DeleteClothById(cloth.ID)
	if err != nil {
		return deletedCloth, err
	}

	return deletedCloth, nil
}

func (s *service) CreateClothImage(input CreateClothImageInput, fileLocation string) (ClothImage, error) {
	cloth, err := s.repository.FindClothByID(input.ClothID)
	if err != nil {
		return ClothImage{}, err
	}

	if cloth.UserID != input.User.ID {
		return ClothImage{}, errors.New("not an owner of the cloth")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.ClothID)
		if err != nil {
			return ClothImage{}, err
		}
	}

	clothImage := ClothImage{}
	clothImage.ClothID = input.ClothID
	clothImage.IsPrimary = isPrimary
	clothImage.FileName = fileLocation

	newClothImage, err := s.repository.CreateClothImage(clothImage)
	if err != nil {
		return newClothImage, err
	}

	return newClothImage, nil
}
