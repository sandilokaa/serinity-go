package handler

import (
	"cheggstore/helper"
	"cheggstore/material"
	"cheggstore/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type materialHandler struct {
	service material.Service
}

func NewMaterialHandler(service material.Service) *materialHandler {
	return &materialHandler{service}
}

func (h *materialHandler) GetAllMaterial(c *gin.Context) {

	search := c.Query("search")

	materials, err := h.service.GetAllMaterial(search)
	if err != nil {
		response := helper.APIResponse("Failed to get material", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get material", http.StatusOK, "success", material.FormatMaterials(materials))
	c.JSON(http.StatusOK, response)
}

func (h *materialHandler) GetMaterialById(c *gin.Context) {
	var input material.GetMaterialDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get material", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	material, err := h.service.GetMaterialById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get materialn", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get material", http.StatusOK, "success", material.FormatMaterialDetail(material))
	c.JSON(http.StatusOK, response)
}

func (h *materialHandler) CreateMaterial(c *gin.Context) {
	var input material.CreateMaterialInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create material", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newMaterial, err := h.service.CreateMaterial(input)
	if err != nil {
		response := helper.APIResponse("Failed to create material", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create material", http.StatusOK, "success", material.FormatMaterial(newMaterial))
	c.JSON(http.StatusOK, response)
}

func (h *materialHandler) UpdateMaterial(c *gin.Context) {
	var inputID material.GetMaterialDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update material", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData material.CreateMaterialInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update material", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldMaterial, err := h.service.GetMaterialById(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to find material", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedMaterial, err := h.service.UpdateMaterial(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update material", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update material", http.StatusOK, "success", material.UpdatedFormatMaterial(updatedMaterial, oldMaterial))
	c.JSON(http.StatusOK, response)

}

func (h *materialHandler) DeleteMaterial(c *gin.Context) {
	var inputID material.GetMaterialDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete material", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedMaterial, err := h.service.DeleteMaterial(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete material", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete material", http.StatusOK, "success", deletedMaterial)
	c.JSON(http.StatusOK, response)
}
