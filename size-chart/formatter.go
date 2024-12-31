package sizechart

type SizeChartFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FileName string `json:"file_name"`
}

func FormatSizeChart(sizeChart SizeChart) SizeChartFormatter {
	sizeChartFormatter := SizeChartFormatter{}
	sizeChartFormatter.ID = sizeChart.ID
	sizeChartFormatter.Name = sizeChart.Name
	sizeChartFormatter.FileName = sizeChart.FileName

	return sizeChartFormatter
}

func UpdatedFormatSizeChart(updatedSizeChart SizeChart, oldSizeChart SizeChart) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldSizeChart.Name != updatedSizeChart.Name {
		updatedFields["name"] = updatedSizeChart.Name
	}

	if oldSizeChart.FileName != updatedSizeChart.FileName {
		updatedFields["file_name"] = updatedSizeChart.FileName
	}

	return updatedFields
}
