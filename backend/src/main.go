package main

import (
	"log"
	"os"
)

func main() {
	cfg := LoadConfig()
	if err := os.MkdirAll(cfg.UploadDir, 0o755); err != nil {
		log.Fatalf("cannot create upload dir: %v", err)
	}

	db := ConnectDB(cfg)
	defer db.Close()

	handler := NewHandler(NewRepository(db), cfg)
	router := NewRouter(handler, cfg)

	log.Printf("listening on :%s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))
}
