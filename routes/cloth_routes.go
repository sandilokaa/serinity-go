package routes

import (
	"serinitystore/cloth"
	"serinitystore/handler"
	"serinitystore/middleware"

	"github.com/gin-gonic/gin"
)

func clothRoutes(router *gin.RouterGroup, clothService cloth.Service) {
	clothHandler := handler.NewClothHandler(clothService)

	router.GET("/items", clothHandler.FindAllCloth)
	router.GET("/items/:id", clothHandler.FindClothByID)

	adminRoutes := router.Group("/items", middleware.RoleMiddleware("admin"))
	{
		adminRoutes.POST("", clothHandler.SaveCloth)
		adminRoutes.PUT("/:id", clothHandler.UpdateClothByID)
		adminRoutes.PUT("/variation/:id", clothHandler.UpdateClothVariationByID)
		adminRoutes.DELETE("/:id", clothHandler.DeleteClothByID)
		adminRoutes.POST("/upload-images", clothHandler.UploadImage)
	}
}
