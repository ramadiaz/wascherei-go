package helpers

import (
	"net/http"
	"wascherei-go/api/users/dto"
	"wascherei-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

func GetUserData(ctx *gin.Context) (dto.UserOutput, *exceptions.Exception){
	var result dto.UserOutput
	user_data, _ := ctx.Get("user")

	result, ok := user_data.(dto.UserOutput)
	if !ok {
		return result, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidTokenStructure)
	}
	

	return result, nil
}