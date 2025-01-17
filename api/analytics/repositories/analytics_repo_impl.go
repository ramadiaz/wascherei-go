package repositories

import (
	"time"
	"wascherei-go/api/analytics/dto"
	"wascherei-go/models"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) ThisMonthIncome(ctx *gin.Context, tx *gorm.DB, userUUID string) (dto.MonthlyIncome, *exceptions.Exception) {
	var income dto.MonthlyIncome

	currentTime := time.Now()
	income.Year = currentTime.Year()
	income.Month = int(currentTime.Month())

	startOfMonth := time.Date(income.Year, time.Month(income.Month), 1, 0, 0, 0, 0, currentTime.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	var result dto.IncomeResult

	err := tx.Model(&models.Transactions{}).
		Where("user_uuid = ? AND created_at >= ? AND created_at < ?", userUUID, startOfMonth, endOfMonth).
		Select("SUM(total_price) as total_income, COUNT(*) as transaction_count").
		Scan(&result).Error

	if err != nil {
		return income, exceptions.ParseGormError(err)
	}

	daysInMonth := currentTime.Sub(startOfMonth).Hours() / 24

	income.Total = int(result.TotalIncome)
	if daysInMonth > 0 {
		income.Average = int(result.TotalIncome / int64(daysInMonth))
	} else {
		income.Average = 0
	}

	return income, nil
}
