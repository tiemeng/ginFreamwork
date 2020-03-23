package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReJson(ctx *gin.Context,data interface{}){
	ctx.JSON(http.StatusOK,data)

}
