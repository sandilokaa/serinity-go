package routes

import (
	"serinitystore/handler"
	"serinitystore/middleware"
	"serinitystore/supplier"

	"github.com/gin-gonic/gin"
)

func supplierRoutes(router *gin.RouterGroup, supplierService supplier.Service) {
	supplierHandler := handler.NewSupplierHandler(supplierService)

	adminRoutes := router.Group("/suppliers", middleware.RoleMiddleware("admin"))
	{
		adminRoutes.GET("", supplierHandler.FindAllSupplier)
		adminRoutes.GET("/:id", supplierHandler.FindSupplierByID)
		adminRoutes.POST("", supplierHandler.CreateSupplier)
		adminRoutes.PUT("/:id", supplierHandler.UpdateSupplierByID)
		adminRoutes.DELETE("/:id", supplierHandler.DeleteSupplierByID)
	}
}
