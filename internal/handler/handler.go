package handler

import (
	"net/http"

	_ "github.com/ShekleinAleksey/project-aurora/docs"
	"github.com/ShekleinAleksey/project-aurora/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Handler struct {
	PurchaseHandler *PurchaseHandler
	CategoryHandler *CategoryHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		PurchaseHandler: NewPurchaseHandler(s.PurchaseService),
		CategoryHandler: NewCategoryHandler(s.CategoryService),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	// Создаём объект Prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router) // регистрация middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{
		categories := api.Group("/categories")
		{
			categories.GET("", h.CategoryHandler.GetAllCategories)
			categories.POST("", h.CategoryHandler.CreateCategory)
			categories.GET("/:id", h.CategoryHandler.GetCategory)
			categories.PUT("/:id", h.CategoryHandler.UpdateCategory)
			categories.DELETE("/:id", h.CategoryHandler.DeleteCategory)
		}
	}

	return router
}
