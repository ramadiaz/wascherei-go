package services

import (
	"wascherei-go/internal/auth/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Login(ctx *gin.Context, data dto.Login) (*string, *exceptions.Exception)
}