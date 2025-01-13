package routers

import (
	testController "wascherei-go/api/test/controllers"
	"wascherei-go/injectors"
	"wascherei-go/pkg/config"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CompRouters(api *gin.RouterGroup) {
	db := config.InitDB()
	validate := validator.New(validator.WithRequiredStructEnabled())

	api.Use(middleware.ClientTracker(db))
	api.Use(middleware.GzipResponseMiddleware())

	api.GET("/ping", testController.Ping)

	userController := injectors.InitializeUserController(db, validate)

	AuthRoutes(api, userController)
}
