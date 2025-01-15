package services

import (
	"wascherei-go/api/products/dto"
	"wascherei-go/api/products/repositories"
	"wascherei-go/pkg/exceptions"
	"wascherei-go/pkg/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.ProductInput) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapProductInputToModel(data)
	input.UUID = uuid.NewString()

	err := s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	return nil
}