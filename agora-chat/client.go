/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"net/http"
	"time"
)

type Client struct {
	appConfig   *App
	appToken    string
	appTokenExp uint32
	httpClient  *http.Client
}

func NewClient(appConfig *App) (*Client, error) {
	client := &Client{
		appConfig:  appConfig,
		httpClient: &http.Client{},
	}
	client.appTokenExp = uint32(time.Now().Unix()) + (24 * 60 * 60)

	appToken, err := client.Tokens().generateChatAppToken()
	if err != nil {
		return nil, err
	}
	client.appToken = appToken
	return client, nil
}

func (c *Client) Tokens() *TokenManager {
	return &TokenManager{c}
}

func (c *Client) Push() *PushManager {
	return &PushManager{c}
}
