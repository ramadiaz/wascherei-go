package routers

import (
	"wascherei-go/api/products/controllers"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, productController controllers.CompControllers) {
	productGroup := r.Group("/product")
	productGroup.Use(middleware.AuthMiddleware())
	{
		productGroup.POST("/create", productController.Create)
		productGroup.GET("/getall", productController.FindByUserUUID)
	}
}
