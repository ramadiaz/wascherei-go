package services

import (
	"wascherei-go/api/users/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Login(ctx *gin.Context, data dto.Login) (*string, *exceptions.Exception)
}