/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	"time"

	"github.com/CarlsonYuan/agora-chat-cli/http"
)

type Client struct {
	appConfig   *App
	appToken    string
	appTokenExp uint32
	pushClient  http.Client[pushResponseResult]
}

func NewClient() *Client {
	client := &Client{
		appConfig:   GetActiveApp(),
		pushClient:  http.NewClient[pushResponseResult](),
		appTokenExp: uint32(time.Now().Unix()) + (24 * 60 * 60),
	}
	appToken, err := client.Tokens().generateChatAppToken()
	if err != nil {
		fmt.Printf("error generate app token")
	}
	client.appToken = appToken
	return client
}

func (c *Client) Tokens() *TokenManager {
	return &TokenManager{c}
}

func (c *Client) Push() *PushManager {
	return &PushManager{c}
}

func (c *Client) Device() *DeviceManager {
	return &DeviceManager{c}
}
