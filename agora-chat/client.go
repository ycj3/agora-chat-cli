/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	"time"

	"github.com/ycj3/agora-chat-cli/http"
)

//go:generate go run go.uber.org/mock/mockgen -source=client.go -destination=client_mock.go -package=agora_chat
type Client interface {
	Tokens() *TokenManager
	User() *UserManager
	Push() *PushManager
	Provider() *ProviderManager
	Device() *DeviceManager
	Message() *MessageManager
}

type client struct {
	appConfig      *App
	appToken       string
	appTokenExp    uint32
	messageClient  http.Client[messageResponseResult]
	userClient     http.Client[userResponseResult]
	pushClient     http.Client[PushResponseResult]
	providerClient http.Client[PrividerResponseResult]
	deviceClient   http.Client[deviceResponseResult]
}

func NewClient() (Client, error) {

	app, err := GetActiveApp()
	if app == nil {
		return nil, fmt.Errorf("Failed to get active app: %w", err)
	}

	client := &client{
		appConfig:      app,
		messageClient:  http.NewClient[messageResponseResult](),
		userClient:     http.NewClient[userResponseResult](),
		pushClient:     http.NewClient[PushResponseResult](),
		providerClient: http.NewClient[PrividerResponseResult](),
		deviceClient:   http.NewClient[deviceResponseResult](),
		appTokenExp:    uint32(time.Hour.Seconds() * 24),
	}
	appToken, err := client.Tokens().GenerateChatAppToken()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate app token")
	}
	client.appToken = appToken
	return client, nil
}

func (c *client) Tokens() *TokenManager {
	return &TokenManager{c}
}

func (c *client) User() *UserManager {
	return &UserManager{c}
}

func (c *client) Push() *PushManager {
	return &PushManager{c}
}

func (c *client) Provider() *ProviderManager {
	return &ProviderManager{c}
}

func (c *client) Device() *DeviceManager {
	return &DeviceManager{c}
}

func (c *client) Message() *MessageManager {
	return &MessageManager{c}
}
