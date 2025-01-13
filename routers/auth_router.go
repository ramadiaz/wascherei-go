package routers

import (
	"wascherei-go/api/users/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", userController.Login)
	}
}
