package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// ConnectDB подключается к Postgres с ожиданием готовности БД.
func ConnectDB(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 30; i++ {
		if err = db.Ping(); err == nil {
			db.SetMaxOpenConns(20)
			return db
		}
		log.Printf("waiting for db: %v", err)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("db is not reachable")
	return nil
}
