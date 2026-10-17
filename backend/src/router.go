package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func corsConfig(cfg Config) cors.Config {
	cc := cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
	origins := strings.Split(cfg.CORSOrigins, ",")
	if len(origins) == 1 && strings.TrimSpace(origins[0]) == "*" {
		cc.AllowAllOrigins = true
		return cc
	}
	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}
	cc.AllowOrigins = origins
	return cc
}

// NewRouter регистрирует middleware и все маршруты.
func NewRouter(h *Handler, cfg Config) *gin.Engine {
	gin.SetMode(cfg.GinMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), cors.New(corsConfig(cfg)))

	api := r.Group("/api")
	api.GET("/health", h.Health)
	api.GET("/openapi.json", h.OpenAPI)
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/api/openapi.json")))
	api.Static("/uploads", cfg.UploadDir)

	api.GET("/categories", h.ListCategories)

	api.GET("/menu", h.ListMenu)
	api.GET("/menu/:id", h.GetMenu)
	api.POST("/menu", h.CreateMenu)
	api.PUT("/menu/:id", h.UpdateMenu)
	api.DELETE("/menu/:id", h.DeleteMenu)

	api.POST("/images", h.UploadImage)
	return r
}
