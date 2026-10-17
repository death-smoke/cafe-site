package main

import (
	"os"
	"path/filepath"
)

// Handler хранит зависимости HTTP-обработчиков.
type Handler struct {
	repo *Repository
	cfg  Config
}

func NewHandler(repo *Repository, cfg Config) *Handler {
	return &Handler{repo: repo, cfg: cfg}
}

// cleanupImage удаляет картинку (запись и файл), если она больше не используется.
func (h *Handler) cleanupImage(id int) {
	path, err := h.repo.DeleteImageIfUnused(id)
	if err == nil {
		_ = os.Remove(filepath.Join(h.cfg.UploadDir, filepath.Base(path)))
	}
}
