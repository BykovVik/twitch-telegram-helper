package tgapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Bot struct {
	token string
}

func NewBot(token string) *Bot {
	return &Bot{token: token}
}

func (b *Bot) SendMiniAppButton(chatID int64, appURL string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.token)

	keyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{
					Text: "Открыть Mini App",
					WebApp: &WebAppInfo{
						URL: appURL,
					},
				},
			},
		},
	}

	body := map[string]interface{}{
		"chat_id":      chatID,
		"text":         "Нажмите кнопку ниже для запуска приложения:",
		"reply_markup": keyboard,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = http.Post(endpoint, "application/json", bytes.NewReader(payload))
	return err
}
