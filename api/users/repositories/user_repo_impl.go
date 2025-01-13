package repositories

import (
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

func (r *CompRepositoriesImpl) FindByUsername(ctx *gin.Context, tx *gorm.DB, username string) (*models.Users, *exceptions.Exception) {
	var userData models.Users

	result := tx.Where("username = ?", username).First(&userData)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &userData, nil
}
