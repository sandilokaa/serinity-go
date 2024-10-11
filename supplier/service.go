package supplier

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
	supplier, err := s.repository.FindAllSupplier(search)
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (s *service) FindSupplierByID(input GetSupplierDetailInput) (Supplier, error) {
	supplier, err := s.repository.FindSupplierByID(input.ID)

	if err != nil {
		return supplier, err
	}

	return supplier, nil
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
