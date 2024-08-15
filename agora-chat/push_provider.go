/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"

	"github.com/CarlsonYuan/agora-chat-cli/http"
)

type ProviderManager struct {
	client *Client
}

type PrividerResponseResult struct {
	Response
	Entities []PushProvider `json:"entities"`
}

type APNSConfig struct {
	TeamId string `json:"teamId"`
	KeyId  string `json:"keyId"`
}
type FCMConfig struct {
	PushType  string `json:"pushType"`
	Priority  string `json:"priority"`
	ProjectId string `json:"projectId"`
	Version   string `json:"version"`
}
type HuaweiConfig struct {
	Category      string `json:"category"`
	ActivityClass string `json:"activityClass"`
}

type HonorConfig struct {
	ActivityClass string `json:"activityClass"`
}

type XiaoMiConfig struct {
	ChannelId string `json:"channelId"`
}
type VivoConfig struct {
	Category string `json:"category"`
}
type OppoConfig struct {
	Category string `json:"category"`
}

const (
	PushProviderAPNS   = PushProviderType("APNS")
	PushProviderFCM    = PushProviderType("ANDROID")
	PushProviderHuaWei = PushProviderType("HUAWEIPUSH")
	PushProviderXiaoMi = PushProviderType("XIAOMIPUSH")
	PushProviderVivo   = PushProviderType("VIVOPUSH")
	PushProviderHonor  = PushProviderType("HONOR")
	PushProviderOppo   = PushProviderType("OPPOPUSH")
	PushProviderMeiZu  = PushProviderType("MEIZUPUSH")
)

type PushProviderType = string

const (
	EnvProduct     = EnvironmentType("PRODUCTION")
	EnvDevelopment = EnvironmentType("DEVELOPMENT")
)

type EnvironmentType = string

type PushProvider struct {
	UUID     string           `json:"uuid,omitempty"`
	Type     PushProviderType `json:"provider"`
	Name     string           `json:"name"`
	Created  int64            `json:"created,omitempty"`
	Modified int64            `json:"modified,omitempty"`
	Disabled bool             `json:"disabled,omitempty"`

	Env         EnvironmentType `json:"environment,omitempty"`
	PackageName string          `json:"packageName,omitempty"`
	Certificate string          `json:"certificate,omitempty"`

	ApnsPushSettings   *APNSConfig   `json:"apnsPushSettings,omitempty"`
	FcmPushSettings    *FCMConfig    `json:"googlePushSettings,omitempty"`
	HuaweiPushSettings *HuaweiConfig `json:"huaweiPushSettings,omitempty"`
	XiaomiPushSetings  *XiaoMiConfig `json:"xiaomiPushSetings,omitempty"`
	VivoPushSettings   *VivoConfig   `json:"vivoPushSettings,omitempty"`
	HonorPushSettings  *HonorConfig  `json:"honorPushSettings,omitempty"`
	OppoPushSettings   *OppoConfig   `json:"oppoPushSettings,omitempty"`
}

// // InsertPushProvider inserts a push provider.
// func (c *Client) InsertPushProvider(ctx context.Context, provider *PushProvider) (*PushProviderListResponse, error) {
// 	body := provider
// 	var resp PushProviderListResponse
// 	err := c.makeRequest(ctx, http.MethodPost, "notifiers", nil, body, &resp)
// 	return &resp, err
// }

// // DeletePushProvider deletes a push provider by uuid.
// func (c *Client) DeletePushProvider(ctx context.Context, uuid string) (*PushProviderListResponse, error) {
// 	var resp PushProviderListResponse
// 	p := path.Join("notifiers", url.PathEscape(uuid))
// 	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
// 	return &resp, err
// }

// type PushProviderListResponse struct {
// 	Response
// 	PushProviders []PushProvider `json:"entities"`
// }

// ListPushProviders returns the list of push providers.
func (pp *ProviderManager) ListPushProviders() (PrividerResponseResult, error) {
	req := pp.ListPushProvidersRequest()
	res, err := pp.client.providerClient.Send(req)
	if err != nil {
		return PrividerResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	return res.Data, err
}
func (pp *ProviderManager) ListPushProvidersRequest() http.Request {

	return http.Request{
		URL:            pp.providerURL(),
		Method:         http.MethodGET,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + pp.client.appToken,
		},
	}
}

func (pp *ProviderManager) providerURL() string {
	baseURL := pp.client.appConfig.BaseURL
	return fmt.Sprintf(baseURL + "/notifiers")
}
