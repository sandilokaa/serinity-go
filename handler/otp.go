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

	emailBody := fmt.Sprintf(
		`
		<div style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 70px 200px;">
			<div style="padding: 5px 40px; background-color: #000000; border-radius: 5px 5px 0 0; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
				<p style="font-size: 30px; color: #FFFFFF;">Permintaan merubah kata sandi</p>
			</div>
			<div style="background-color: #FFFFFF; border-radius: 0 0 5px 5px; padding: 20px 40px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); font-size: 14px;">
				<p>Halo %s,</p>
				<p>Seseorang telah meminta kata sandi baru untuk akun berikut di Serinity:</p>
				<p>Jika Anda tidak membuat permintaan, abaikan saja email ini. Jika Anda ingin melanjutkan:</p>
				<p style="font-size: 18px !important; font-weight: bold; color: #0000000;">Kode OTP Anda: %s</p>
				<p style="color: #888888;">Kode ini berlaku hingga: %s</p>
			</div>
			<div style="margin-top: 20px;">
				<p style="font-size: 12px; text-align: center;">Serinity — Built with Love</p>
			</div>
		</div>
		`,
		input.Email,
		input.Otp,
		input.Expiry.Format(time.RFC3339),
	)
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
