package routes

import (
	"serinitystore/handler"
	"serinitystore/middleware"
	sizechart "serinitystore/size-chart"

	"github.com/gin-gonic/gin"
)

func sizeChartRoutes(router *gin.RouterGroup, sizeChartService sizechart.Service) {
	sizeChartHandler := handler.NewSizeChartHandler(sizeChartService)

	adminRoutes := router.Group("/size-charts", middleware.RoleMiddleware("admin"))
	{
		adminRoutes.POST("", sizeChartHandler.SaveSizeChart)
		adminRoutes.PUT("/:id", sizeChartHandler.UpdateSizeChart)
	}
}
