package repositories

import (
	"wascherei-go/models"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Example) *exceptions.Exception
}
