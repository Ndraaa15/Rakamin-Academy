package message

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func Error(ctx *gin.Context, statusCode int64, message string, data interface{}) {
	ctx.JSON(int(statusCode), Response{
		Message: message,
		Code:    statusCode,
		Status:  false,
		Data:    data,
	})
}

func Success(ctx *gin.Context, statusCode int64, message string, data interface{}) {
	ctx.JSON(int(statusCode), Response{
		Message: message,
		Code:    statusCode,
		Status:  true,
		Data:    data,
	})
}
