package tgapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (b *Bot) GetUpdates(offset int) ([]Update, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", b.token, offset)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result struct {
		OK     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	err = json.Unmarshal(raw, &result)

	if err != nil {
		return nil, err
	}

	return result.Result, nil
}
