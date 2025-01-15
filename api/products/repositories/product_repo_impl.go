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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) (*[]models.Products, *exceptions.Exception) {
	var products []models.Products

	result := tx.Find(&products)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &products, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, UUID string) (*models.Products, *exceptions.Exception) {
	var product models.Products

	result := tx.Where("uuid = ?", UUID).First(&product)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &product, nil
}

func (r *CompRepositoriesImpl) FindByUserUUID(ctx *gin.Context, tx *gorm.DB, UserUUID string) (*[]models.Products, *exceptions.Exception) {
	var products []models.Products

	result := tx.Where("user_uuid = ?", UserUUID).Find(&products)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &products, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception {
	result := tx.Save(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception {
	result := tx.Delete(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}
