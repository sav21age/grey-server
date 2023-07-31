package response

import (
	"github.com/gin-gonic/gin"
)

type MsgWrapper struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, MsgWrapper{message})
}