/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	"time"

	auth "github.com/ycj3/agora-chat-cli/agora-chat/auth"
	"github.com/ycj3/agora-chat-cli/http"
)

//go:generate go run go.uber.org/mock/mockgen -source=client.go -destination=client_mock.go -package=agora_chat
type Client interface {
	User() *UserManager
	Push() *PushManager
	Provider() *ProviderManager
	Device() *DeviceManager
	Message() *MessageManager
}

type client struct {
	appConfig      *App
	appToken       string
	messageClient  http.Client[messageResponseResult]
	userClient     http.Client[userResponseResult]
	pushClient     http.Client[PushResponseResult]
	providerClient http.Client[PrividerResponseResult]
	deviceClient   http.Client[deviceResponseResult]
}

func NewClient() (Client, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %s", err)
	}

	app, err := cfg.GetActiveApp()
	if app == nil {
		return nil, fmt.Errorf("failed to get active app: %s", err)
	}

	client := &client{
		appConfig:      app,
		messageClient:  http.NewClient[messageResponseResult](),
		userClient:     http.NewClient[userResponseResult](),
		pushClient:     http.NewClient[PushResponseResult](),
		providerClient: http.NewClient[PrividerResponseResult](),
		deviceClient:   http.NewClient[deviceResponseResult](),
	}

	expire := uint32(time.Hour.Seconds() * 24)
	at, err := auth.NewAuth(app.AppID, app.AppCertificate, expire)

	if err != nil {
		return nil, fmt.Errorf("failed to auth: %s", err)
	}

	appToken, err := at.TokenFromEnvOrBuilder()
	if err != nil {
		return nil, fmt.Errorf("failed to generate app token: %s", err)
	}
	client.appToken = appToken
	return client, nil
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
