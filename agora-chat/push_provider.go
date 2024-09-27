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
	Sound  string `json:"sound"`
}

const (
	FCMPushData         = FCMPushType("DATA")
	FCMPushNotification = FCMPushType("NOTIFICATION")
	FCMPushBoth         = FCMPushType("BOTH")
)

type FCMPushType = string

const (
	FCMPushPriorityHigh   = FCMPushPriorityType("high")
	FCMPushPriorityNormal = FCMPushPriorityType("normal")
)

type FCMPushPriorityType = string

type FCMConfig struct {
	PushType    FCMPushType         `json:"pushType,omitempty"`
	Priority    FCMPushPriorityType `json:"priority,omitempty"`
	ProjectId   string              `json:"projectId,omitempty"`
	Version     string              `json:"version,omitempty"`
	SupportAPNs bool                `json:"supportAPNs"`
	Sound       string              `json:"sound,omitempty"`
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
	UUID       string `json:"uuid,omitempty"`
	NotifierId string `json:"notifierId,omitempty"` // Same as UUID, it is used when editing the notifier.
	Type       string `json:"type,omitempty"`

	Name     string           `json:"name,omitempty"`
	Created  int64            `json:"created,omitempty"`
	Modified int64            `json:"modified,omitempty"`
	Disabled bool             `json:"disabled,omitempty"`
	Provider PushProviderType `json:"provider,omitempty"`

	Env         EnvironmentType `json:"environment,omitempty"`
	PackageName string          `json:"packageName,omitempty"`
	Certificate string          `json:"certificate,omitempty"`
	File        string          `json:"file,omitempty"`
	Passphrase  string          `json:"passphrase,omitempty"`

	ApnsPushSettings   *APNSConfig   `json:"apnsPushSettings,omitempty"`
	FcmPushSettings    *FCMConfig    `json:"googlePushSettings,omitempty"`
	HuaweiPushSettings *HuaweiConfig `json:"huaweiPushSettings,omitempty"`
	XiaomiPushSetings  *XiaoMiConfig `json:"xiaomiPushSetings,omitempty"`
	VivoPushSettings   *VivoConfig   `json:"vivoPushSettings,omitempty"`
	HonorPushSettings  *HonorConfig  `json:"honorPushSettings,omitempty"`
	OppoPushSettings   *OppoConfig   `json:"oppoPushSettings,omitempty"`
}

func (pp *ProviderManager) UpsertPushProvider(provider PushProvider) (PrividerResponseResult, error) {
	req, err := pp.upsertPushProvidersRequest(provider)
	if err != nil {
		return PrividerResponseResult{}, fmt.Errorf("failed to request: %s", err.Error())
	}
	res, err := pp.client.providerClient.Send(req)
	if err != nil {
		return PrividerResponseResult{}, fmt.Errorf("failed sent request: %s", err.Error())
	}
	if res.StatusCode != gohttp.StatusOK {
		return PrividerResponseResult{}, res.Data.Error
	}
	return res.Data, err
}

func (pp *ProviderManager) upsertPushProvidersRequest(p PushProvider) (http.Request, error) {

	if p.NotifierId != "" {
		return http.Request{
			URL:            pp.providerURL(),
			Method:         http.MethodPOST,
			ResponseFormat: http.ResponseFormatJSON,
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": "Bearer " + pp.client.appToken,
			},
			Payload: &http.JSONPayload{Content: p},
		}, nil
	}

	// Build the payload based on the provider type
	content := map[string]interface{}{
		"provider":    p.Provider,
		"name":        p.Name,
		"environment": p.Env,
		"certificate": p.Certificate,
		"packageName": p.PackageName,
	}
	// Add provider-specific settings to the payload
	switch p.Provider {
	case PushProviderAPNS:
		content["teamId"] = p.ApnsPushSettings.TeamId
		content["keyId"] = p.ApnsPushSettings.KeyId
		content["sound"] = p.ApnsPushSettings.Sound
		content["passphrase"] = p.Passphrase
		files := map[string]string{
			"file": p.File,
		}
		return http.Request{
			URL:            pp.providerURL(),
			Method:         http.MethodPOST,
			ResponseFormat: http.ResponseFormatJSON,
			Headers: map[string]string{
				"Authorization": "Bearer " + pp.client.appToken,
			},
			Payload: &http.FormPayload{
				Fields: content,
				Files:  files,
			},
		}, nil
	case PushProviderFCM:
		content["pushType"] = p.FcmPushSettings.PushType
		content["priority"] = p.FcmPushSettings.Priority
		content["supportAPNs"] = p.FcmPushSettings.SupportAPNs
		files := map[string]string{
			"file": p.File,
		}
		return http.Request{
			URL:            pp.providerURL(),
			Method:         http.MethodPOST,
			ResponseFormat: http.ResponseFormatJSON,
			Headers: map[string]string{
				"Authorization": "Bearer " + pp.client.appToken,
			},
			Payload: &http.FormPayload{
				Fields: content,
				Files:  files,
			},
		}, nil
	case PushProviderHuaWei:
		if p.HuaweiPushSettings != nil {
			content["huaweiPushSettings"] = p.HuaweiPushSettings
		}
		// Add cases for other provider types as needed
	default:
		return http.Request{}, fmt.Errorf("unsupport push provider")
	}

	return http.Request{}, nil
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
