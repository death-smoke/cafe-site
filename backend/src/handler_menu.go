package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var sortColumns = map[string]string{
	"id": "m.id", "name": "m.name", "price": "m.price", "category": "c.id",
}

// parseFilter разбирает и валидирует query-параметры списка.
func (h *Handler) parseFilter(c *gin.Context) (MenuFilter, error) {
	f := MenuFilter{Search: strings.TrimSpace(c.Query("search"))}

	if v := c.Query("category_id"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil || id <= 0 {
			return f, errors.New("category_id должен быть положительным числом")
		}
		f.CategoryID = id
	}

	col, ok := sortColumns[c.DefaultQuery("sort", "id")]
	if !ok {
		return f, errors.New("sort: допустимо id, name, price, category")
	}
	f.SortCol = col

	f.Dir = strings.ToUpper(c.DefaultQuery("order", "asc"))
	if f.Dir != "ASC" && f.Dir != "DESC" {
		return f, errors.New("order: допустимо asc или desc")
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		return f, errors.New("page должен быть >= 1")
	}
	f.Page = page

	limit, err := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(h.cfg.DefaultLimit)))
	if err != nil || limit < 1 || limit > h.cfg.MaxLimit {
		return f, fmt.Errorf("limit должен быть от 1 до %d", h.cfg.MaxLimit)
	}
	f.Limit = limit
	return f, nil
}

func (h *Handler) ListMenu(c *gin.Context) {
	f, err := h.parseFilter(c)
	if err != nil {
		fail(c, http.StatusBadRequest, err.Error())
		return
	}
	items, total, err := h.repo.ListMenu(f)
	if err != nil {
		dbError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items, "total": total, "page": f.Page, "limit": f.Limit,
		"total_pages": (total + f.Limit - 1) / f.Limit,
	})
}

func (h *Handler) GetMenu(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	it, err := h.repo.GetMenuItem(id)
	if errors.Is(err, sql.ErrNoRows) {
		fail(c, http.StatusNotFound, "блюдо не найдено")
		return
	}
	if err != nil {
		dbError(c, err)
		return
	}
	c.JSON(http.StatusOK, it)
}

func bindInput(c *gin.Context) (MenuInput, bool) {
	var in MenuInput
	if err := c.ShouldBindJSON(&in); err != nil {
		fail(c, http.StatusBadRequest, "ошибка валидации: проверьте name, category_id, price (>0) и длину текстов")
		return in, false
	}
	in.Name = strings.TrimSpace(in.Name)
	if in.Name == "" {
		fail(c, http.StatusBadRequest, "название не может быть пустым")
		return in, false
	}
	return in, true
}

func (h *Handler) CreateMenu(c *gin.Context) {
	in, ok := bindInput(c)
	if !ok {
		return
	}
	id, err := h.repo.CreateMenuItem(in)
	if err != nil {
		dbError(c, err)
		return
	}
	it, err := h.repo.GetMenuItem(id)
	if err != nil {
		dbError(c, err)
		return
	}
	c.JSON(http.StatusCreated, it)
}

func (h *Handler) UpdateMenu(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	in, ok := bindInput(c)
	if !ok {
		return
	}
	oldImage, err := h.repo.MenuImageID(id)
	if errors.Is(err, sql.ErrNoRows) {
		fail(c, http.StatusNotFound, "блюдо не найдено")
		return
	}
	if err != nil {
		dbError(c, err)
		return
	}
	if err := h.repo.UpdateMenuItem(id, in); err != nil {
		dbError(c, err)
		return
	}
	if oldImage != nil && (in.ImageID == nil || *in.ImageID != *oldImage) {
		h.cleanupImage(*oldImage)
	}
	it, err := h.repo.GetMenuItem(id)
	if err != nil {
		dbError(c, err)
		return
	}
	c.JSON(http.StatusOK, it)
}

func (h *Handler) DeleteMenu(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	imageID, err := h.repo.DeleteMenuItem(id)
	if errors.Is(err, sql.ErrNoRows) {
		fail(c, http.StatusNotFound, "блюдо не найдено")
		return
	}
	if err != nil {
		dbError(c, err)
		return
	}
	if imageID != nil {
		h.cleanupImage(*imageID)
	}
	c.Status(http.StatusNoContent)
}
