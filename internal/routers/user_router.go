package routers

import (
	"wascherei-go/api/users/controllers"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	userGroup := r.Group("/user")
	userGroup.Use(middleware.InternalMiddleware())
	{
		userGroup.POST("/create", userController.Create)
	}
}
