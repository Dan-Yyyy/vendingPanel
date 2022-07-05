package handler

import (
	_ "github.com/Dan-Yyyy/vendingPanel.git/docs"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
