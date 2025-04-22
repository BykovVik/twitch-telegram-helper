package handlers

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func StatusHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Its WORK")
	}
}
