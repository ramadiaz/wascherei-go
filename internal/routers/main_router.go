package routers

import (
	"wascherei-go/injectors"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InternalRouters(r *gin.RouterGroup) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	
	r.Use(middleware.GzipResponseMiddleware())
	internalController := injectors.InitializeInternalAuthController(validate)

	AuthRoutes(r, internalController)
}
