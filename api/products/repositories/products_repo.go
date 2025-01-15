package repositories

import (
	"wascherei-go/models"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) (*[]models.Products, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, tx *gorm.DB, UUID string) (*models.Products, *exceptions.Exception)
	FindByUserUUID(ctx *gin.Context, tx *gorm.DB, UserUUID string) (*[]models.Products, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception
}
