package services

import (
	"wascherei-go/api/analytics/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	ThisMonthIncome(ctx *gin.Context, userUUID string) (dto.MonthlyIncome, *exceptions.Exception)
}