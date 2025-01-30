package routes

import (
	"serinitystore/handler"
	"serinitystore/middleware"
	"serinitystore/transaction"

	"github.com/gin-gonic/gin"
)

func transactionRoutes(router *gin.RouterGroup, transactionService transaction.Service) {
	transactionHandler := handler.NewTransactionHandler(transactionService)

	userRoutes := router.Group("/items/transactions/user")
	{
		userRoutes.POST("", transactionHandler.CreateTransaction)
		userRoutes.GET("/:userId", transactionHandler.GetTransactionByUserID)
		userRoutes.GET("/:userId/:id", transactionHandler.GetTransactionUserIDByID)
	}

	adminRoutes := router.Group("/items/transactions", middleware.RoleMiddleware("admin"))
	{
		adminRoutes.GET("", transactionHandler.FindAllTransaction)
		adminRoutes.GET("/:id", transactionHandler.GetTransactionByID)
	}

	router.POST("/transactions/notification", transactionHandler.GetNotification)
}
