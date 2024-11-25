package category

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
	categories, err := s.repository.FindAllCategory(search)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *service) FindCategoryByID(input GetCategoryDetailInput) (Category, error) {
	category, err := s.repository.FindCategoryByID(input.ID)

	if err != nil {
		return category, err
	}

	return category, nil
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
