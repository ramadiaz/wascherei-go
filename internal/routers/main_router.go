package routers

import (
	"wascherei-go/injectors"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	r.Use(middleware.GzipResponseMiddleware())
	internalController := injectors.InitializeInternalAuthController(validate)
	userController := injectors.InitializeUserController(db, validate)

	AuthRoutes(r, internalController)
	UserRoutes(r, userController)
}
