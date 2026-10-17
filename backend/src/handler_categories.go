package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListCategories(c *gin.Context) {
	cats, err := h.repo.ListCategories()
	if err != nil {
		dbError(c, err)
		return
	}
	c.JSON(http.StatusOK, cats)
}
