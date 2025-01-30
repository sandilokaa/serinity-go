package routes

import (
	"serinitystore/handler"
	"serinitystore/material"
	"serinitystore/middleware"

	"github.com/gin-gonic/gin"
)

func materialRoutes(router *gin.RouterGroup, materialService material.Service) {
	materialHandler := handler.NewMaterialHandler(materialService)

	adminRoutes := router.Group("/materials", middleware.RoleMiddleware("admin"))
	{
		adminRoutes.GET("", materialHandler.GetAllMaterial)
		adminRoutes.GET("/:id", materialHandler.GetMaterialById)
		adminRoutes.POST("", materialHandler.CreateMaterial)
		adminRoutes.PUT("/:id", materialHandler.UpdateMaterial)
		adminRoutes.DELETE("/:id", materialHandler.DeleteMaterial)
	}
}
