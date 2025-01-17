package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	ThisMonthIncome(ctx *gin.Context)
}