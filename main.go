package main

import (
	"cheggstore/auth"
	"cheggstore/cloth"
	"cheggstore/handler"
	"cheggstore/material"
	"cheggstore/middleware"
	"cheggstore/supplier"
	"cheggstore/user"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/chegg_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	materialRepository := material.NewRepository(db)
	supplierRepository := supplier.NewRepository(db)
	clothRepository := cloth.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	materialService := material.NewService(materialRepository)
	supplierService := supplier.NewService(supplierRepository)
	clothService := cloth.NewService(clothRepository)

	userHandler := handler.NewHandler(userService, authService)
	materialHandler := handler.NewMaterialHandler(materialService)
	supplierHandler := handler.NewSupplierHandler(supplierService)
	clothHandler := handler.NewClothHandler(clothService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.GET("/materials", materialHandler.GetAllMaterial)
	api.GET("/materials/:id", materialHandler.GetMaterialById)
	api.GET("/suppliers", supplierHandler.FindAllSupplier)
	api.GET("/suppliers/:id", supplierHandler.FindSupplierByID)
	api.GET("/cloths", clothHandler.FindAllCloth)
	api.GET("/cloths/:id", clothHandler.FindClothByID)

	protectedRoutes := api.Group("/protected", middleware.AuthMiddleware(authService, userService))
	{
		protectedRoutes.GET("/auth/me", userHandler.CurrentUser)
		protectedRoutes.POST("/materials", materialHandler.CreateMaterial)
		protectedRoutes.PUT("/materials/:id", materialHandler.UpdateMaterial)
		protectedRoutes.DELETE("/materials/:id", materialHandler.DeleteMaterial)
		protectedRoutes.POST("/suppliers", supplierHandler.CreateSupplier)
		protectedRoutes.PUT("/suppliers/:id", supplierHandler.UpdateSupplierByID)
		protectedRoutes.DELETE("/suppliers/:id", supplierHandler.DeleteSupplierByID)
		protectedRoutes.POST("/cloths", clothHandler.SaveCloth)
		protectedRoutes.PUT("/cloths/:id", clothHandler.UpdateClothByID)
		protectedRoutes.DELETE("/cloths/:id", clothHandler.DeleteClothByID)
		protectedRoutes.POST("/cloths/upload-images", clothHandler.UploadImage)
		protectedRoutes.PUT("/cloths/upload-images/:id", clothHandler.UpdateClothImage)
	}

	router.Run()
}
