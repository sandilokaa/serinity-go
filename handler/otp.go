package handler

import (
	"fmt"
	"net/http"
	"serinitystore/helper"
	"serinitystore/otp"
	"time"

	"github.com/gin-gonic/gin"
)

type otpHandler struct {
	service otp.Service
}

func NewOTPHandler(service otp.Service) *otpHandler {
	return &otpHandler{service}
}

func (h *otpHandler) SaveOTP(c *gin.Context) {
	var input otp.OTPRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Otp request failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.Otp = helper.GenerateOTP()
	input.Expiry = time.Now().Add(3 * time.Minute)

	err = h.service.SaveOTP(input)
	if err != nil {
		response := helper.APIResponse("Otp request failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	emailBody := fmt.Sprintf("Kode OTP Anda adalah: %s\nKode ini berlaku hingga: %s", input.Otp, input.Expiry.Format(time.RFC3339))
	err = helper.SendEmail(input.Email, "Kode OTP Anda", emailBody)
	if err != nil {
		response := helper.APIResponse("Failed to send email", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("OTP has been sent to your email", http.StatusOK, "success", input)
	c.JSON(http.StatusOK, response)
}

func (h *otpHandler) VerifyOTP(c *gin.Context) {
	var input otp.VerifyOTPInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get OTP", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isValid, err := h.service.VerifyOTP(input.Email, input.Otp)
	if err != nil {
		response := helper.APIResponse("Failed to verify OTP", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !isValid {
		response := helper.APIResponse("Invalid OTP", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("OTP verified successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
