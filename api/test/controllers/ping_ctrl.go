package ping

import (
	"net/http"
	"wascherei-go/api/test/dto"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "pong",
	})
}