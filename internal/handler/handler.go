package handler

import (
	"github.com/akram620/alif/internal/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Events
}

func NewHandler(service service.Events) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(logger.SetLogger())
	router.Use(gin.Recovery())

	router.GET("/health", h.Health)

	ap1V1 := router.Group("/api/v1")
	{
		ap1V1.GET("/test", h.Test)
	}

	return router
}
