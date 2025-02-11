package services

import (
	"wascherei-go/api/transactions/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.TransactionInput) *exceptions.Exception
}