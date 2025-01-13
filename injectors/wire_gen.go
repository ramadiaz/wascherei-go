// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	controllers2 "wascherei-go/api/users/controllers"
	"wascherei-go/api/users/repositories"
	services2 "wascherei-go/api/users/services"
	"wascherei-go/internal/auth/controllers"
	"wascherei-go/internal/auth/services"
)

// Injectors from injector.go:

func InitializeInternalAuthController(validate *validator.Validate) controllers.CompControllers {
	compServices := services.NewComponentServices(validate)
	compControllers := controllers.NewCompController(compServices)
	return compControllers
}

func InitializeUserController(db *gorm.DB, validate *validator.Validate) controllers2.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	compServices := services2.NewComponentServices(compRepositories, db, validate)
	compControllers := controllers2.NewCompController(compServices)
	return compControllers
}

// injector.go:

var internalAuthFeatureSet = wire.NewSet(services.NewComponentServices, controllers.NewCompController)

var userFeatureSet = wire.NewSet(repositories.NewComponentRepository, services2.NewComponentServices, controllers2.NewCompController)
