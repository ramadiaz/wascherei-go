package services

import (
	"wascherei-go/api/analytics/dto"
	"wascherei-go/api/analytics/repositories"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (s *CompServicesImpl) ThisMonthIncome(ctx *gin.Context, userUUID string) (dto.MonthlyIncome, *exceptions.Exception) {
	return s.repo.ThisMonthIncome(ctx, s.DB, userUUID)
}