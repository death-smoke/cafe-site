package main

import (
	_ "embed"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed openapi.json
var openapiSpec []byte

func (h *Handler) Health(c *gin.Context) {
	if err := h.repo.Ping(); err != nil {
		fail(c, http.StatusServiceUnavailable, "db unavailable")
		return
	}
	host, _ := os.Hostname()
	c.JSON(http.StatusOK, gin.H{"status": "ok", "host": host})
}

func (h *Handler) OpenAPI(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", openapiSpec)
}
