package routes

import (
	"serinitystore/auth"
	"serinitystore/category"
	"serinitystore/cloth"
	"serinitystore/material"
	"serinitystore/middleware"
	"serinitystore/otp"
	sizechart "serinitystore/size-chart"
	"serinitystore/supplier"
	"serinitystore/transaction"
	"serinitystore/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	authService auth.Service,
	userService user.Service,
	materialService material.Service,
	categoryService category.Service,
	supplierService supplier.Service,
	sizeChartService sizechart.Service,
	clothService cloth.Service,
	transactionService transaction.Service,
	otpService otp.Service,
) {
	api := router.Group("/api/v1")

	userRoutes(api, userService, authService)
	clothRoutes(api, clothService)
	transactionRoutes(api, transactionService)
	otpRoutes(api, otpService)

	protected := api.Group("/protected", middleware.AuthMiddleware(authService, userService))
	{
		materialRoutes(protected, materialService)
		supplierRoutes(protected, supplierService)
		categoryRoutes(protected, categoryService)
		sizeChartRoutes(protected, sizeChartService)
		clothRoutes(protected, clothService)
		transactionRoutes(protected, transactionService)
	}
}
