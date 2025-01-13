// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	userControllers "wascherei-go/api/users/controllers"
	userRepositories "wascherei-go/api/users/repositories"
	userServices "wascherei-go/api/users/services"

	internalAuthControllers "wascherei-go/internal/auth/controllers"
	internalAuthServices "wascherei-go/internal/auth/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var internalAuthFeatureSet = wire.NewSet(
	internalAuthServices.NewComponentServices,
	internalAuthControllers.NewCompController,
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	userServices.NewComponentServices,
	userControllers.NewCompController,
)

func InitializeInternalAuthController(validate *validator.Validate) internalAuthControllers.CompControllers {
	wire.Build(internalAuthFeatureSet)
	return nil
}

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}