package main

import (
	"cheggstore/auth"
	"cheggstore/cloth"
	"cheggstore/handler"
	"cheggstore/material"
	"cheggstore/middleware"
	"cheggstore/payment"
	"cheggstore/supplier"
	"cheggstore/transaction"
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
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	materialService := material.NewService(materialRepository)
	supplierService := supplier.NewService(supplierRepository)
	clothService := cloth.NewService(clothRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, clothRepository, paymentService)

	userHandler := handler.NewHandler(userService, authService)
	materialHandler := handler.NewMaterialHandler(materialService)
	supplierHandler := handler.NewSupplierHandler(supplierService)
	clothHandler := handler.NewClothHandler(clothService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.GET("/cloths", clothHandler.FindAllCloth)
	api.GET("/cloths/:id", clothHandler.FindClothByID)

	protectedRoutes := api.Group("/protected", middleware.AuthMiddleware(authService, userService))
	{
		protectedRoutes.GET("/auth/me", userHandler.CurrentUser)
		protectedRoutes.GET("/materials", middleware.RoleMiddleware("admin"), materialHandler.GetAllMaterial)
		protectedRoutes.GET("/materials/:id", middleware.RoleMiddleware("admin"), materialHandler.GetMaterialById)
		protectedRoutes.POST("/materials", middleware.RoleMiddleware("admin"), materialHandler.CreateMaterial)
		protectedRoutes.PUT("/materials/:id", middleware.RoleMiddleware("admin"), materialHandler.UpdateMaterial)
		protectedRoutes.DELETE("/materials/:id", middleware.RoleMiddleware("admin"), materialHandler.DeleteMaterial)
		protectedRoutes.GET("/suppliers", middleware.RoleMiddleware("admin"), supplierHandler.FindAllSupplier)
		protectedRoutes.GET("/suppliers/:id", middleware.RoleMiddleware("admin"), supplierHandler.FindSupplierByID)
		protectedRoutes.POST("/suppliers", middleware.RoleMiddleware("admin"), supplierHandler.CreateSupplier)
		protectedRoutes.PUT("/suppliers/:id", middleware.RoleMiddleware("admin"), supplierHandler.UpdateSupplierByID)
		protectedRoutes.DELETE("/suppliers/:id", middleware.RoleMiddleware("admin"), supplierHandler.DeleteSupplierByID)
		protectedRoutes.POST("/cloths", middleware.RoleMiddleware("admin"), clothHandler.SaveCloth)
		protectedRoutes.PUT("/cloths/:id", middleware.RoleMiddleware("admin"), clothHandler.UpdateClothByID)
		protectedRoutes.PUT("/cloths/variation/:id", middleware.RoleMiddleware("admin"), clothHandler.UpdateClothVariationByID)
		protectedRoutes.DELETE("/cloths/:id", middleware.RoleMiddleware("admin"), clothHandler.DeleteClothByID)
		protectedRoutes.POST("/cloths/upload-images", middleware.RoleMiddleware("admin"), clothHandler.UploadImage)
		protectedRoutes.POST("/cloths/transactions", transactionHandler.CreateTransaction)
		protectedRoutes.GET("/cloths/:userId/transactions", transactionHandler.GetTransactionByUserID)
		protectedRoutes.GET("/cloths/:userId/transactions/:id", transactionHandler.GetTransactionUserIDByID)
		protectedRoutes.GET("/cloths/transactions", middleware.RoleMiddleware("admin"), transactionHandler.FindAllTransaction)
		protectedRoutes.GET("/cloths/transactions/:id", middleware.RoleMiddleware("admin"), transactionHandler.GetTransactionByID)
		protectedRoutes.POST("/transactions/notification", transactionHandler.GetNotification)
	}

	router.Run()
}
