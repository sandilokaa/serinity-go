package sizechart

type Service interface {
	SaveSizeChart(input CreateSizeChartInput, fileLocation string) (SizeChart, error)
	UpdateSizeChartByID(inputID SizeChartInputDetail, inputData UpdateSizeChartInput, fileLocation string) (SizeChart, error)
	FindSizeChartByID(input SizeChartInputDetail) (SizeChart, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveSizeChart(input CreateSizeChartInput, fileLocation string) (SizeChart, error) {
	sizeChart := SizeChart{}
	sizeChart.Name = input.Name
	sizeChart.FileName = fileLocation
	sizeChart.UserID = input.User.ID

	newSizeChart, err := s.repository.SaveSizeChart(sizeChart)
	if err != nil {
		return newSizeChart, err
	}

	return newSizeChart, nil
}

func (s *service) FindSizeChartByID(input SizeChartInputDetail) (SizeChart, error) {
	sizeChart, err := s.repository.FindSizeChartByID(input.ID)
	if err != nil {
		return sizeChart, err
	}

	return sizeChart, nil
}

func (s *service) UpdateSizeChartByID(inputID SizeChartInputDetail, inputData UpdateSizeChartInput, fileLocation string) (SizeChart, error) {

	sizeChart, err := s.repository.FindSizeChartByID(inputID.ID)
	if err != nil {
		return sizeChart, err
	}

	if inputData.Name != "" {
		sizeChart.Name = inputData.Name
	}

	if fileLocation != "" {
		if sizeChart.FileName != "" {
			err := s.repository.DeleteImage(sizeChart.FileName)
			if err != nil {
				return sizeChart, err
			}
		}

		sizeChart.FileName = fileLocation
	}

	updatedSizeChart, err := s.repository.UpdateSizeChartByID(sizeChart)
	if err != nil {
		return updatedSizeChart, err
	}

	return updatedSizeChart, nil
}
