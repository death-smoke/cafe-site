package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadImage принимает PNG (multipart: file, alt), сохраняет файл и запись в images.
func (h *Handler) UploadImage(c *gin.Context) {
	max := h.cfg.MaxUploadBytes
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, max+(1<<20))

	fh, err := c.FormFile("file")
	if err != nil {
		var mbe *http.MaxBytesError
		if errors.As(err, &mbe) {
			fail(c, http.StatusRequestEntityTooLarge, "файл слишком большой")
			return
		}
		fail(c, http.StatusBadRequest, "поле file обязательно (multipart/form-data)")
		return
	}
	if fh.Size > max {
		fail(c, http.StatusRequestEntityTooLarge, fmt.Sprintf("максимальный размер файла: %d МБ", max>>20))
		return
	}

	f, err := fh.Open()
	if err != nil {
		fail(c, http.StatusBadRequest, "не удалось прочитать файл")
		return
	}
	defer f.Close()

	cfg, err := png.DecodeConfig(f)
	if err != nil {
		fail(c, http.StatusBadRequest, "допустим только формат PNG")
		return
	}

	b := make([]byte, 16)
	_, _ = rand.Read(b)
	name := hex.EncodeToString(b) + ".png"
	if err := c.SaveUploadedFile(fh, filepath.Join(h.cfg.UploadDir, name)); err != nil {
		log.Printf("save file: %v", err)
		fail(c, http.StatusInternalServerError, "не удалось сохранить файл")
		return
	}

	attrs := map[string]any{
		"filename":      name,
		"original_name": filepath.Base(fh.Filename),
		"mime_type":     "image/png",
		"size_bytes":    fh.Size,
		"width":         cfg.Width,
		"height":        cfg.Height,
		"alt":           strings.TrimSpace(c.PostForm("alt")),
		"uploaded_at":   time.Now().UTC().Format(time.RFC3339),
	}
	raw, _ := json.Marshal(attrs)
	path := "uploads/" + name

	id, err := h.repo.InsertImage(path, raw)
	if err != nil {
		_ = os.Remove(filepath.Join(h.cfg.UploadDir, name))
		dbError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id, "path": path, "url": "/api/" + path, "attributes": attrs})
}
