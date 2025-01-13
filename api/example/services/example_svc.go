package services

import (
	"wascherei-go/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompService interface {
	Create(ctx *gin.Context, data dto.ExampleInput) *exceptions.Exception
}
