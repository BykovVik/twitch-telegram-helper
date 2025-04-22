package main

import (
	"log"
	"net/http"
	"twich-telegram-helper/handlers"
	"twich-telegram-helper/internal/config"
	"twich-telegram-helper/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cfg := config.Load()
	db := storage.InitDB(cfg.DatabaseUrl)

	http.HandleFunc("/api/status", handlers.StatusHandler(db))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
