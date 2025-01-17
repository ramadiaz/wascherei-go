package controllers

import (
	"net/http"
	"wascherei-go/api/analytics/dto"
	"wascherei-go/api/analytics/services"
	"wascherei-go/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) ThisMonthIncome(ctx *gin.Context) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	data, err := h.services.ThisMonthIncome(ctx, userData.UUID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "find success",
	})
}
