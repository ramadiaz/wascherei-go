package controllers

import (
	"net/http"
	"wascherei-go/api/transactions/dto"
	"wascherei-go/api/transactions/services"
	"wascherei-go/pkg/exceptions"
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

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.TransactionInput

	errRequest := ctx.ShouldBindJSON(&data)
	if errRequest != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, errRequest.Error()))
		return
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	data.UserUUID = userData.UUID
	err = h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "create success",
	})
}
