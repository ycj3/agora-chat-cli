/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"

	"github.com/CarlsonYuan/agora-chat-cli/http"
)

type MessageManager struct {
	client *Client
}

type MessageType string

const (
	MessageTypeText   MessageType = "txt"    // Text message
	MessageTypeImage  MessageType = "img"    // Image message
	MessageTypeAudio  MessageType = "audio"  // Voice message
	MessageTypeVideo  MessageType = "video"  // Video message
	MessageTypeFile   MessageType = "file"   // File message
	MessageTypeLoc    MessageType = "loc"    // Location message
	MessageTypeCmd    MessageType = "cmd"    // Command message
	MessageTypeCustom MessageType = "custom" // Custom message
)

type messageResponseResult struct {
	Response
	Path            string            `json:"path"`
	URI             string            `json:"uri"`
	Organization    string            `json:"organization"`
	Application     string            `json:"application"`
	Data            map[string]string `json:"data"`
	ApplicationName string            `json:"applicationName"`
}

func (mm *MessageManager) SendUserMessage(from string, to []string, msgType MessageType, body map[string]interface{}, options map[string]interface{}) (map[string]string, error) {
	request := mm.sendUserMessageRequest(from, to, msgType, body, options)
	res, err := mm.client.messageClient.Send(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Data, nil
}

func (mm *MessageManager) sendUserMessageRequest(from string, to []string, msgType MessageType, body map[string]interface{}, options map[string]interface{}) http.Request {

	var reqBody = map[string]interface{}{
		"from": from,
		"to":   to,
		"type": msgType,
		"body": body,
	}

	for key, value := range options {
		reqBody[key] = value
	}

	req := http.Request{
		URL:            mm.sendUserMessageURL(),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + mm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: reqBody,
		},
	}
	return req
}

func (mm *MessageManager) sendUserMessageURL() string {
	baseURL := mm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/messages/users", baseURL)
}

func (mm *MessageManager) sendChatgroupMessageURL() string {
	baseURL := mm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/messages/chatgroups", baseURL)
}

func (mm *MessageManager) sendChatroomMessageURL() string {
	baseURL := mm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/messages/chatrooms", baseURL)
}
