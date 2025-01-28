package handler

import (
	"net/http"
	"serinitystore/helper"
	"serinitystore/supplier"
	"serinitystore/user"

	"github.com/gin-gonic/gin"
)

type supplierHandler struct {
	service supplier.Service
}

func NewSupplierHandler(service supplier.Service) *supplierHandler {
	return &supplierHandler{service}
}

func (h *supplierHandler) CreateSupplier(c *gin.Context) {
	var input supplier.CreateSupplierInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newSupplier, err := h.service.CreateSupplier(input)
	if err != nil {
		response := helper.APIResponse("Failed to create supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create supplier", http.StatusOK, "success", supplier.FormatSupplier(newSupplier))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) FindAllSupplier(c *gin.Context) {

	search := c.Query("search")

	suppliers, err := h.service.FindAllSupplier(search)
	if err != nil {
		response := helper.APIResponse("Failed to find supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find supplier", http.StatusOK, "success", supplier.FormatSuppliers(suppliers))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) FindSupplierByID(c *gin.Context) {
	var input supplier.GetSupplierDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to find supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	supplier, err := h.service.FindSupplierByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to find supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find supplier", http.StatusOK, "success", supplier.FormatSupplierDetail(supplier))
	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) UpdateSupplierByID(c *gin.Context) {
	var inputID supplier.GetSupplierDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData supplier.UpdateSupplierInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldSupllier, err := h.service.FindSupplierByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to find supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedSupplier, err := h.service.UpdateSupplierByID(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update supplier", http.StatusOK, "success", supplier.UpdatedFormatSupplier(updatedSupplier, oldSupllier))
	c.JSON(http.StatusOK, response)

}

func (h *supplierHandler) DeleteSupplierByID(c *gin.Context) {
	var inputID supplier.GetSupplierDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedSupplier, err := h.service.DeleteSupplierByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete supplier", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete supplier", http.StatusOK, "success", deletedSupplier)
	c.JSON(http.StatusOK, response)
}
