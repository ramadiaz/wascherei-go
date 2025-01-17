package services

import (
	"wascherei-go/api/transactions/dto"
	"wascherei-go/api/transactions/repositories"
	"wascherei-go/pkg/exceptions"
	"wascherei-go/pkg/mapper"

	productRepo "wascherei-go/api/products/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo        repositories.CompRepositories
	productRepo productRepo.CompRepositories
	DB          *gorm.DB
	validate    *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, productRepo productRepo.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:        compRepositories,
		productRepo: productRepo,
		DB:          db,
		validate:    validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.TransactionInput) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}
	
	productData, err := s.productRepo.FindByUUID(ctx, s.DB, data.ProductUUID)
	if err != nil {
		return err
	}

	input := mapper.MapTransactionInputToModel(data)
	input.UUID = uuid.NewString()

	input.TotalPrice = uint(float32(productData.Price) * input.UnitSize)

	err = s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	return nil
}
