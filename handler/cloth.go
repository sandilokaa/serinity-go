package handler

import (
	"cheggstore/cloth"
	"cheggstore/helper"
	"cheggstore/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clothHandler struct {
	service cloth.Service
}

func NewClothHandler(service cloth.Service) *clothHandler {
	return &clothHandler{service}
}

func (h *clothHandler) SaveCloth(c *gin.Context) {
	var input cloth.CreateClothInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create cloth", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCloth, err := h.service.SaveCloth(input)
	if err != nil {
		response := helper.APIResponse("Failed to create cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create cloth", http.StatusOK, "success", cloth.FormatCloth(newCloth))
	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) FindAllCloth(c *gin.Context) {

	search := c.Query("search")

	cloths, err := h.service.FindAllCloth(search)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find cloth", http.StatusOK, "success", cloth.FormatCloths(cloths))
	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) FindClothByID(c *gin.Context) {
	var input cloth.ClothInputDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cloth, err := h.service.FindClothByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find cloth", http.StatusOK, "success", cloth.FormatClothDetail(cloth))
	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) UpdateClothByID(c *gin.Context) {
	var inputID cloth.ClothInputDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get cloth", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData cloth.UpdateClothInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update cloth", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldCloth, err := h.service.FindClothByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCloth, err := h.service.UpdateClothByID(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to updated cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated cloth", http.StatusOK, "success", cloth.UpdatedFormatCloth(updatedCloth, oldCloth))
	c.JSON(http.StatusOK, response)

}

func (h *clothHandler) DeleteClothByID(c *gin.Context) {
	var inputID cloth.ClothInputDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedCloth, err := h.service.DeleteClothByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to deleted cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to deleted cloth", http.StatusOK, "success", deletedCloth)
	c.JSON(http.StatusOK, response)

}

func (h *clothHandler) UploadImage(c *gin.Context) {
	var input cloth.CreateClothImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload cloth image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload cloth image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload cloth image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.CreateClothImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload cloth image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success to upload cloth image", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) UpdateClothImage(c *gin.Context) {
	var inputID cloth.ClothImageInputDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get cloth image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData cloth.UpdateClothImageInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update cloth image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload cloth image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload cloth image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	oldClothImage, err := h.service.FindClothImageByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedClothImage, err := h.service.UpdateClothImage(inputID, inputData, path)
	if err != nil {
		response := helper.APIResponse("Failed to updated cloth image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated cloth image", http.StatusOK, "success", cloth.UpdateFormatClothImage(updatedClothImage, oldClothImage))
	c.JSON(http.StatusOK, response)

}
