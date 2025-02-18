package routes

import (
	"serinitystore/handler"
	"serinitystore/otp"

	"github.com/gin-gonic/gin"
)

func otpRoutes(router *gin.RouterGroup, otpService otp.Service) {
	otpHandler := handler.NewOTPHandler(otpService)

	otpRoutes := router.Group("/forgot-password")
	{
		otpRoutes.POST("/send-email", otpHandler.SaveOTP)
	}
}
