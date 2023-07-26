package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	JsonResponse(c, http.StatusOK, 200, "OK", data)
}

func JsonErrorResponse(c *gin.Context, code int, msg string) {
	JsonResponse(c, http.StatusInternalServerError, code, msg, nil)
}

func JsonInternalServerErrorResponse(c *gin.Context) {
	JsonResponse(c, http.StatusInternalServerError, 200500, "Internal server error", nil)
}
