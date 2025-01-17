package routers

import (
	"wascherei-go/api/analytics/controllers"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AnalyticsRoutes(r *gin.RouterGroup, analyticController controllers.CompControllers) {
	analyticGroup := r.Group("/analytic")
	analyticGroup.Use(middleware.AuthMiddleware())
	{
		currentGroup := analyticGroup.Group("/current")
		{
			currentGroup.GET("/month", analyticController.ThisMonthIncome)
		}
	}
}
