package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func fail(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}

// dbError переводит ошибки Postgres в HTTP-ответы.
func dbError(c *gin.Context, err error) {
	var pe *pq.Error
	if errors.As(err, &pe) {
		switch pe.Code {
		case "23505":
			fail(c, http.StatusConflict, "блюдо с таким названием уже существует")
			return
		case "23503":
			fail(c, http.StatusBadRequest, "указанная категория или изображение не существует")
			return
		}
	}
	log.Printf("db error: %v", err)
	fail(c, http.StatusInternalServerError, "внутренняя ошибка сервера")
}

func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		fail(c, http.StatusBadRequest, "некорректный id")
		return 0, false
	}
	return id, true
}
