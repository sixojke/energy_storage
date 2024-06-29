package v1

import (
	"energy_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

type dataResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
