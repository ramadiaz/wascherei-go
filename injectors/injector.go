// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	internalAuthControllers "wascherei-go/internal/auth/controllers"
	internalAuthServices "wascherei-go/internal/auth/services"

	userControllers "wascherei-go/api/users/controllers"
	userRepositories "wascherei-go/api/users/repositories"
	userServices "wascherei-go/api/users/services"

	productControllers "wascherei-go/api/products/controllers"
	productRepositories "wascherei-go/api/products/repositories"
	productServices "wascherei-go/api/products/services"
	
	transactionControllers "wascherei-go/api/transactions/controllers"
	transactionRepositories "wascherei-go/api/transactions/repositories"
	transactionServices "wascherei-go/api/transactions/services"

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

var productFeatureSet = wire.NewSet(
	productRepositories.NewComponentRepository,
	productServices.NewComponentServices,
	productControllers.NewCompController,
)

var transactionFeatureSet = wire.NewSet(
	transactionRepositories.NewComponentRepository,
	transactionServices.NewComponentServices,
	transactionControllers.NewCompController,
)

func InitializeInternalAuthController(validate *validator.Validate) internalAuthControllers.CompControllers {
	wire.Build(internalAuthFeatureSet)
	return nil
}

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeProductController(db *gorm.DB, validate *validator.Validate) productControllers.CompControllers {
	wire.Build(productFeatureSet)
	return nil
}

func InitializeTransactionController(db *gorm.DB, validate *validator.Validate) transactionControllers.CompControllers {
	wire.Build(transactionFeatureSet)
	return nil
}
