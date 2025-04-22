package tgapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type Bot struct {
	token string
}

func NewBot(token string) *Bot {
	return &Bot{token: token}
}

func (b *Bot) SendMessage(chatID int64, text string) {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.token)
	http.PostForm(endpoint, url.Values{
		"chat_id": {fmt.Sprintf("%d", chatID)},
		"text":    {text},
	})
}
