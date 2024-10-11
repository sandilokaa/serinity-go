package material

type Service interface {
	GetAllMaterial(search string) ([]Material, error)
	GetMaterialById(input GetMaterialDetailInput) (Material, error)
	CreateMaterial(input CreateMaterialInput) (Material, error)
	UpdateMaterial(inputID GetMaterialDetailInput, inputData CreateMaterialInput) (Material, error)
	DeleteMaterial(inputID GetMaterialDetailInput) (Material, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllMaterial(search string) ([]Material, error) {
	material, err := s.repository.FindAllMaterial(search)
	if err != nil {
		return material, err
	}

	return material, nil
}

func (s *service) GetMaterialById(input GetMaterialDetailInput) (Material, error) {
	material, err := s.repository.FindMaterialById(input.ID)

	if err != nil {
		return material, err
	}

	return material, nil
}

func (s *service) CreateMaterial(input CreateMaterialInput) (Material, error) {
	material := Material{}
	material.UserID = input.User.ID
	material.MaterialName = input.MaterialName

	newMaterial, err := s.repository.SaveMaterial(material)
	if err != nil {
		return newMaterial, err
	}

	return newMaterial, nil
}

func (s *service) UpdateMaterial(inputID GetMaterialDetailInput, inputData CreateMaterialInput) (Material, error) {
	material, err := s.repository.FindMaterialById(inputID.ID)
	if err != nil {
		return material, err
	}

	material.MaterialName = inputData.MaterialName

	updatedMaterial, err := s.repository.UpdateMaterial(material)
	if err != nil {
		return updatedMaterial, err
	}

	return updatedMaterial, nil
}

func (s *service) DeleteMaterial(inputID GetMaterialDetailInput) (Material, error) {
	material, err := s.repository.FindMaterialById(inputID.ID)
	if err != nil {
		return material, err
	}

	deletedMaterial, err := s.repository.DeleteMaterial(material.ID)
	if err != nil {
		return deletedMaterial, err
	}

	return deletedMaterial, nil
}
