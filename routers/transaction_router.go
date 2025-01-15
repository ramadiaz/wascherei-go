package routers

import (
	"wascherei-go/api/transactions/controllers"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup, transactionController controllers.CompControllers) {
	transactionGroup := r.Group("/transaction")
	transactionGroup.Use(middleware.AuthMiddleware())
	{
		transactionGroup.POST("/create", transactionController.Create)
	}
}
