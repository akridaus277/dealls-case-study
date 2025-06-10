package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBadRequestResponse(message string) Response {
	return Response{
		Status:  http.StatusText(http.StatusBadRequest),
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    nil,
	}
}

func NewInternalServerErrorResponse(message string) Response {
	return Response{
		Status:  http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}

func NewNotFoundResponse(message string) Response {
	return Response{
		Status:  http.StatusText(http.StatusNotFound),
		Code:    http.StatusNotFound,
		Message: message,
		Data:    nil,
	}
}

func NewUnauthorizedResponse(message string) Response {
	return Response{
		Status:  http.StatusText(http.StatusUnauthorized),
		Code:    http.StatusUnauthorized,
		Message: message,
		Data:    nil,
	}
}

func NewForbiddenResponse(message string) Response {
	return Response{
		Status:  http.StatusText(http.StatusForbidden),
		Code:    http.StatusForbidden,
		Message: message,
		Data:    nil,
	}
}

func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}

func SendResponse(c *gin.Context, response Response) {
	c.JSON(response.Code, response)
}
