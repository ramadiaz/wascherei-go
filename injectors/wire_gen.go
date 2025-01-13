// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"wascherei-go/api/example/controllers"
	"wascherei-go/api/example/repositories"
	"wascherei-go/api/example/services"
)

// Injectors from injector.go:

func InitializeExampleController(db *gorm.DB, validate *validator.Validate) controllers.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	compService := services.NewComponentServices(compRepositories, db, validate)
	compControllers := controllers.NewCompController(compService)
	return compControllers
}

// injector.go:

var exampleFeatureSet = wire.NewSet(repositories.NewComponentRepository, services.NewComponentServices, controllers.NewCompController)
