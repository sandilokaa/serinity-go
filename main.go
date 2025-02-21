package main

import (
	"log"
	"serinitystore/auth"
	"serinitystore/category"
	"serinitystore/cloth"
	"serinitystore/material"
	"serinitystore/otp"
	"serinitystore/payment"
	"serinitystore/redis"
	"serinitystore/routes"
	sizechart "serinitystore/size-chart"
	"serinitystore/supplier"
	"serinitystore/transaction"
	"serinitystore/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dsn := "root:@tcp(127.0.0.1:3306)/serinity_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Redis connection
	redis.InitRedis()

	// Repository initialization
	userRepository := user.NewRepository(db)
	materialRepository := material.NewRepository(db)
	supplierRepository := supplier.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	sizeChartRepository := sizechart.NewRepository(db)
	clothRepository := cloth.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	otpRepository := otp.NewOTPRepository(redis.RedisClient)

	// Service initialization
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	materialService := material.NewService(materialRepository)
	supplierService := supplier.NewService(supplierRepository)
	categoryService := category.NewService(categoryRepository)
	sizeChartService := sizechart.NewService(sizeChartRepository)
	clothService := cloth.NewService(clothRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, clothRepository, paymentService)
	otpService := otp.NewOTPService(otpRepository)

	// Router setup
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")

	routes.RegisterRoutes(
		router,
		authService,
		userService,
		materialService,
		categoryService,
		supplierService,
		sizeChartService,
		clothService,
		transactionService,
		otpService,
	)

	router.Run()
}
