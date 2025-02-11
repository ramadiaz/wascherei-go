package controllers

import (
	"net/http"
	"wascherei-go/api/products/dto"
	"wascherei-go/api/products/services"
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
	var data dto.ProductInput

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

func (h *CompControllersImpl) FindByUserUUID(ctx *gin.Context) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	data, err := h.services.FindByUserUUID(ctx, userData.UUID)
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