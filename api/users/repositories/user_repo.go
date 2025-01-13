package repositories

import (
	"wascherei-go/models"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	FindByUsername(ctx *gin.Context, tx *gorm.DB, username string) (*models.Users, *exceptions.Exception)
}