package helpers

import "github.com/gin-gonic/gin"

type FormatResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FormatCollectionResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Response(context *gin.Context, code int, message string, data interface{}) {
	response := FormatResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	context.JSON(code, response)
}

func SuccessResponse(context *gin.Context, code int, message string, data interface{}) {
	Response(context, code, message, data)
}

func ErrorResponse(context *gin.Context, code int, message string, data interface{}) {
	Response(context, code, message, data)
}
