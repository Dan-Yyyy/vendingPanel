package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type responseError struct {
	Message string `json:"message"`
}

type responseStatus struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, responseMessage string) {
	log.Printf(message)
	c.AbortWithStatusJSON(statusCode, responseError{responseMessage})
}
