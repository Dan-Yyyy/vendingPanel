package handler

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn)
	}

	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
