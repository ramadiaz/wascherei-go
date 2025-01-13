package routers

import (
	"net/http"
	"wascherei-go/dto"
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

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.Response{
			Status: http.StatusOK,
			Message:   "pong",
		})
	})

	exampleController := injectors.InitializeExampleController(db, validate)

	ExampleRoutes(api, exampleController)
}
