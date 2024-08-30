/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	gohttp "net/http"

	"github.com/ycj3/agora-chat-cli/http"
)

type ProviderManager struct {
	client *client
}

type PrividerResponseResult struct {
	Error
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

func (pp *ProviderManager) InsertPushProvider(provider PushProvider) (PrividerResponseResult, error) {
	req := pp.instertPushProvidersRequest(provider)
	res, err := pp.client.providerClient.Send(req)
	if err != nil {
		return PrividerResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	if res.StatusCode != gohttp.StatusOK {
		return PrividerResponseResult{}, res.Data.Error
	}
	return res.Data, err
}

func (pp *ProviderManager) instertPushProvidersRequest(provider PushProvider) http.Request {
	// Build the payload based on the provider type
	content := map[string]interface{}{
		"provider":    provider.Type,
		"name":        provider.Name,
		"environment": provider.Env,
		"certificate": provider.Certificate,
		"packageName": provider.PackageName,
	}

	// Add provider-specific settings to the payload
	switch provider.Type {
	case PushProviderAPNS:
		if provider.ApnsPushSettings != nil {
			content["apnsPushSettings"] = provider.ApnsPushSettings
		}
	case PushProviderFCM:
		if provider.FcmPushSettings != nil {
			content["googlePushSettings"] = provider.FcmPushSettings
		}
	case PushProviderHuaWei:
		if provider.HuaweiPushSettings != nil {
			content["huaweiPushSettings"] = provider.HuaweiPushSettings
		}
		// Add cases for other provider types as needed
	}

	return http.Request{
		URL:            pp.providerURL(),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + pp.client.appToken,
		},
		Payload: &http.JSONPayload{Content: content},
	}
}

// DeletePushProvider deletes a push provider by uuid.
func (pp *ProviderManager) DeletePushProvider(uuid string) (PrividerResponseResult, error) {
	req := pp.deletePushProviderRequest(uuid)
	res, err := pp.client.providerClient.Send(req)
	if err != nil {
		return PrividerResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	if res.StatusCode != gohttp.StatusOK {
		return PrividerResponseResult{}, res.Data.Error
	}
	return res.Data, err
}

func (pp *ProviderManager) deletePushProviderRequest(uuid string) http.Request {
	return http.Request{
		URL:            pp.providerURL() + "/" + uuid,
		Method:         http.MethodDELETE,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + pp.client.appToken,
		},
	}
}

// ListPushProviders returns the list of push providers.
func (pp *ProviderManager) ListPushProviders() (PrividerResponseResult, error) {
	req := pp.listPushProvidersRequest()
	res, err := pp.client.providerClient.Send(req)
	if err != nil {
		return PrividerResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	if res.StatusCode != gohttp.StatusOK {
		return PrividerResponseResult{}, res.Data.Error
	}
	return res.Data, err
}

func (pp *ProviderManager) listPushProvidersRequest() http.Request {

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
