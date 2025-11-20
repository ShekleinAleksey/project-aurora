package handler

import (
	"github.com/ShekleinAleksey/project-aurora/internal/service"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Handler struct {
	PurchaseHandler *PurchaseHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		PurchaseHandler: NewPurchaseHandler(s.PurchaseService),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Создаём объект Prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router) // регистрация middleware
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{
		categories := api.Group("/categories")
		{
			categories.GET("/")
			categories.POST("/")
		}
	}

	return router
}
