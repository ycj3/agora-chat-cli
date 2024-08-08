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
	appConfig     *App
	appToken      string
	appTokenExp   uint32
	messageClient http.Client[messageResponseResult]
	userClient    http.Client[userResponseResult]
	pushClient    http.Client[pushResponseResult]
	deviceClient  http.Client[deviceResponseResult]
}

func NewClient() *Client {
	client := &Client{
		appConfig:     GetActiveApp(),
		messageClient: http.NewClient[messageResponseResult](),
		userClient:    http.NewClient[userResponseResult](),
		pushClient:    http.NewClient[pushResponseResult](),
		deviceClient:  http.NewClient[deviceResponseResult](),
		appTokenExp:   uint32(time.Hour.Seconds() * 24),
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

func (c *Client) User() *UserManager {
	return &UserManager{c}
}

func (c *Client) Push() *PushManager {
	return &PushManager{c}
}

func (c *Client) Device() *DeviceManager {
	return &DeviceManager{c}
}

func (c *Client) Message() *MessageManager {
	return &MessageManager{c}
}
