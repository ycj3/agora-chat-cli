/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"

	"github.com/CarlsonYuan/agora-chat-cli/http"
)

type PushManager struct {
	client *Client
}

type pushResult struct {
	PushStatus string                 `json:"pushStatus"`
	Data       map[string]interface{} `json:"data,omitempty"`
	Desc       *string                `json:"desc,omitempty"`
}
type PushResponseResult struct {
	Response
	Data []pushResult `json:"data"`
}

type pushStrategy int

const (
	PushPrividerFirstThenAgoraChat pushStrategy = iota // 0
	OnlyAgoraChat                                      // 1
	OnlyPushPrivider                                   // 2 (Default)
	AgoraFirstThenPushPrivider                         // 3
	OnlyOnlineAgoraChat                                // 4
)

type PushMessage struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	SubTitle string `json:"sub_title"`
}

func (pm *PushManager) SyncPush(userID string, strategy pushStrategy, msg PushMessage) (PushResponseResult, error) {

	request := pm.syncPushRequest(userID, strategy, msg)
	res, err := pm.client.pushClient.Send(request)
	if err != nil {
		return PushResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	return res.Data, nil
}

func (pm *PushManager) syncPushRequest(userID string, strategy pushStrategy, msg PushMessage) http.Request {

	return http.Request{
		URL:            pm.syncPushURL(userID),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + pm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: map[string]interface{}{
				"strategy":    strategy,
				"pushMessage": msg,
			},
		},
	}
}

func (pm *PushManager) syncPushURL(userID string) string {
	baseURL := pm.client.appConfig.BaseURL
	return fmt.Sprintf(baseURL+"/push/sync/"+"%s", userID)
}
