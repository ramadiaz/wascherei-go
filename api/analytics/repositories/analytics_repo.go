package repositories

import (
	"wascherei-go/api/analytics/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	ThisMonthIncome(ctx *gin.Context, tx *gorm.DB, userUUID string) (dto.MonthlyIncome, *exceptions.Exception)
}