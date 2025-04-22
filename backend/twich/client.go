package twitch

import (
	"log"

	"github.com/gempir/go-twitch-irc/v4"
)

type Client struct {
	client      *twitch.Client
	channelName string
}

func NewClient(username, token, channel string) *Client {
	c := twitch.NewClient(username, "oauth:"+token)
	return &Client{client: c, channelName: channel}
}

func (c *Client) StartListening(callback func(msg string)) {
	c.client.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		callback(msg.Message)
	})
	c.client.Join(c.channelName)
	if err := c.client.Connect(); err != nil {
		log.Fatal(err)
	}
}
