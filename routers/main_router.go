package routers

import (
	testController "wascherei-go/api/test/controllers"
	"wascherei-go/injectors"
	"wascherei-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(api *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	api.Use(middleware.ClientTracker(db))
	api.Use(middleware.GzipResponseMiddleware())

	api.GET("/ping", testController.Ping)

	userController := injectors.InitializeUserController(db, validate)
	productController := injectors.InitializeProductController(db, validate)
	transactionController := injectors.InitializeTransactionController(db, validate)

	AuthRoutes(api, userController)
	ProductRoutes(api, productController)
	TransactionRoutes(api, transactionController)
}
