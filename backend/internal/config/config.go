package config

import "os"

type Config struct {
	TelegramToken string
	TwichToken    string
	TwichUser     string
	ChannelName   string
	DatabaseUrl   string
}

func Load() Config {
	return Config{
		TelegramToken: os.Getenv("TG_API_TOKEN"),
		TwichToken:    os.Getenv(""),
		TwichUser:     os.Getenv(""),
		ChannelName:   os.Getenv(""),
		DatabaseUrl:   os.Getenv("DATABASE_URL"),
	}
}
