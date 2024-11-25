package handler

import (
	"cheggstore/category"
	"cheggstore/helper"
	"cheggstore/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service category.Service
}

func NewCategoryHandler(service category.Service) *categoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CreateCategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCategory, err := h.service.CreateCategory(input)
	if err != nil {
		response := helper.APIResponse("Failed to create category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create category", http.StatusOK, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) FindAllCategory(c *gin.Context) {

	search := c.Query("search")

	categories, err := h.service.FindAllCategory(search)
	if err != nil {
		response := helper.APIResponse("Failed to find category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find category", http.StatusOK, "success", category.FormatCategories(categories))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) FindCategoryByID(c *gin.Context) {
	var input category.GetCategoryDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to find category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, err := h.service.FindCategoryByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to find category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find category", http.StatusOK, "success", category.FormatCategoryDetail(category))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdateCategoryByID(c *gin.Context) {
	var inputID category.GetCategoryDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData category.UpdateCategoryInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldCategory, err := h.service.FindCategoryByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to find category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCategory, err := h.service.UpdateCategoryByID(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update category", http.StatusOK, "success", category.UpdatedFormatCategory(updatedCategory, oldCategory))
	c.JSON(http.StatusOK, response)

}

func (h *categoryHandler) DeleteCategoryByID(c *gin.Context) {
	var inputID category.GetCategoryDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedCategory, err := h.service.DeleteCategoryByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete category", http.StatusOK, "success", deletedCategory)
	c.JSON(http.StatusOK, response)
}
