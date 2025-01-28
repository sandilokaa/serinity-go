package handler

import (
	"fmt"
	"net/http"
	"serinitystore/helper"
	sizechart "serinitystore/size-chart"
	"serinitystore/user"

	"github.com/gin-gonic/gin"
)

type sizeChartHandler struct {
	service sizechart.Service
}

func NewSizeChartHandler(service sizechart.Service) *sizeChartHandler {
	return &sizeChartHandler{service}
}

func (h *sizeChartHandler) SaveSizeChart(c *gin.Context) {
	var input sizechart.CreateSizeChartInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload size chart image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload size chart image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload size chart image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newSizeChart, err := h.service.SaveSizeChart(input, path)
	if err != nil {
		response := helper.APIResponse("Failed to create size chart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create size chart", http.StatusOK, "success", sizechart.FormatSizeChart(newSizeChart))
	c.JSON(http.StatusOK, response)
}

func (h *sizeChartHandler) UpdateSizeChart(c *gin.Context) {
	var inputID sizechart.SizeChartInputDetail
	var inputData sizechart.UpdateSizeChartInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to get size chart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil && err.Error() != "http: no such file" {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload size chart image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser
	userID := currentUser.ID

	var path string
	if file != nil {
		path = fmt.Sprintf("images/%d-%s", userID, file.Filename)

		err = c.SaveUploadedFile(file, path)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.APIResponse("Failed to upload size chart image", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update size chart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldSizeChart, err := h.service.FindSizeChartByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get size chart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedSizeChart, err := h.service.UpdateSizeChartByID(inputID, inputData, path)
	if err != nil {
		response := helper.APIResponse("Failed to update size chart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update size chart", http.StatusOK, "success", sizechart.UpdatedFormatSizeChart(updatedSizeChart, oldSizeChart))
	c.JSON(http.StatusOK, response)
}
