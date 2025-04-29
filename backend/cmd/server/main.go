package main

import (
	"log"
	"net/http"
	"time"
	"twich-telegram-helper/handlers"
	"twich-telegram-helper/internal/config"
	"twich-telegram-helper/internal/storage"
	"twich-telegram-helper/internal/tgapi"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cfg := config.Load()
	db := storage.InitDB(cfg.DatabaseUrl)

	http.HandleFunc("/api/status", handlers.StatusHandler(db))

	bot := tgapi.NewBot(cfg.TelegramToken)
	StartBotPolling(bot, "dfdfdf")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

}

func StartBotPolling(bot *tgapi.Bot, appURL string) {
	go func() {
		offset := 0
		for {
			updates, err := bot.GetUpdates(offset)
			if err != nil {
				log.Printf("Ошибка при получении обновлений: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}

			for _, update := range updates {
				if update.Message != nil {
					log.Printf("Сообщение: %s", update.Message.Text)
					bot.SendMiniAppButton(update.Message.Chat.ID, appURL)
				}
				offset = update.UpdateID + 1
			}

			time.Sleep(1 * time.Second)
		}
	}()
}
