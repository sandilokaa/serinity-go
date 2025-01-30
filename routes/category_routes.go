package routes

import (
	"serinitystore/category"
	"serinitystore/handler"
	"serinitystore/middleware"

	"github.com/gin-gonic/gin"
)

func categoryRoutes(router *gin.RouterGroup, categoryService category.Service) {
	categoryHandler := handler.NewCategoryHandler(categoryService)

	adminRoutes := router.Group("/categories", middleware.RoleMiddleware("admin"))
	{
		adminRoutes.GET("", categoryHandler.FindAllCategory)
		adminRoutes.GET("/:id", categoryHandler.FindCategoryByID)
		adminRoutes.POST("", categoryHandler.CreateCategory)
		adminRoutes.PUT("/:id", categoryHandler.UpdateCategoryByID)
		adminRoutes.DELETE("/:id", categoryHandler.DeleteCategoryByID)
	}
}
