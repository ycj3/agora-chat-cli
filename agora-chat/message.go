/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"

	"github.com/ycj3/agora-chat-cli/http"
)

type MessageManager struct {
	client *client
}

type messageResponseResult struct {
	Response
	Path            string            `json:"path"`
	URI             string            `json:"uri"`
	Organization    string            `json:"organization"`
	Application     string            `json:"application"`
	Data            map[string]string `json:"data"`
	ApplicationName string            `json:"applicationName"`
}

// SendUserMessage sends a Message to users via Agora Chat Server.
func (mm *MessageManager) SendUsersMessage(message *Message) (map[string]string, error) {
	request, err := mm.sendUsersMessageRequest(message)
	if err != nil {
		return nil, err
	}

	res, err := mm.client.messageClient.Send(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Data, nil
}

func (mm *MessageManager) sendUsersMessageRequest(message *Message) (http.Request, error) {
	if err := validateMessage(message); err != nil {
		return http.Request{}, err
	}

	req := http.Request{
		URL:            mm.sendUsersMessageURL(),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + mm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: message,
		},
	}
	return req, nil
}

func (mm *MessageManager) sendUsersMessageURL() string {
	baseURL := mm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/messages/users", baseURL)
}

// SendGroupMessage sends a Message to groups via Agora Chat Server.
func (mm *MessageManager) SendGroupsMessage(message *Message) (map[string]string, error) {
	request, err := mm.sendGroupsMessageRequest(message)
	if err != nil {
		return nil, err
	}

	res, err := mm.client.messageClient.Send(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Data, nil
}

func (mm *MessageManager) sendGroupsMessageRequest(message *Message) (http.Request, error) {
	if err := validateMessage(message); err != nil {
		return http.Request{}, err
	}

	req := http.Request{
		URL:            mm.sendChatgroupMessageURL(),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + mm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: message,
		},
	}
	return req, nil
}
func (mm *MessageManager) sendChatgroupMessageURL() string {
	baseURL := mm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/messages/chatgroups", baseURL)
}

// SendRoomsMessage sends a Message to rooms via Agora Chat Server.
func (mm *MessageManager) SendRoomsMessage(message *Message) (map[string]string, error) {
	request, err := mm.sendRoomsMessageRequest(message)
	if err != nil {
		return nil, err
	}

	res, err := mm.client.messageClient.Send(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Data, nil
}

func (mm *MessageManager) sendRoomsMessageRequest(message *Message) (http.Request, error) {
	if err := validateMessage(message); err != nil {
		return http.Request{}, err
	}

	req := http.Request{
		URL:            mm.sendChatroomMessageURL(),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + mm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: message,
		},
	}
	return req, nil
}
func (mm *MessageManager) sendChatroomMessageURL() string {
	baseURL := mm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/messages/chatrooms", baseURL)
}
