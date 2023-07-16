package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse[T any] struct {
	Status          string `json:"status"`
	ResponseMessage string `json:"message"`
	Data            T      `json:"data"`
}

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus ResponseStatus, data T) ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) ApiResponse[T] {
	return ApiResponse[T]{
		Status:          status,
		ResponseMessage: message,
		Data:            data,
	}
}

func UnauthorisedResponse(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized,
		BuildResponse_(
			Unauthorized.GetResponseStatus(),
			Unauthorized.GetResponseMessage(),
			"Invalid username or password",
		),
	)
}
