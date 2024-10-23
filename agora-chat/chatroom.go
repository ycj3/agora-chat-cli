/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	gohttp "net/http"
	"strings"

	"github.com/ycj3/agora-chat-cli/http"
)

type chatroomDetail struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	MembersOnly       bool          `json:"membersonly"`
	AllowInvites      bool          `json:"allowinvites"`
	MaxUsers          int           `json:"maxusers"`
	Owner             string        `json:"owner"`
	Created           int64         `json:"created"`
	Custom            string        `json:"custom"`
	Mute              bool          `json:"mute"`
	AffiliationsCount int           `json:"affiliations_count"`
	Avatar            string        `json:"avatar"`
	Disabled          bool          `json:"disabled"`
	Affiliations      []affiliation `json:"affiliations"`
	Public            bool          `json:"public"`
}

type chatroomResponseResult struct {
	Response
	Path            string           `json:"path"`
	URI             string           `json:"uri"`
	Organization    string           `json:"organization"`
	Application     string           `json:"application"`
	Data            []chatroomDetail `json:"data"`
	ApplicationName string           `json:"applicationName"`
	Error
}

type ChatroomManager struct {
	client *client
}

func (gm *ChatroomManager) GetChatroomDetail(rids []string) (chatroomResponseResult, error) {
	request, err := gm.getChatroomDetailRequest(rids)
	if err != nil {
		return chatroomResponseResult{}, err
	}

	res, err := gm.client.chatroomClient.Send(request)

	statusCode := res.StatusCode
	if statusCode != gohttp.StatusOK {
		if statusCode == gohttp.StatusNotFound {
			return chatroomResponseResult{}, res.Data.Error
		}
		return chatroomResponseResult{}, fmt.Errorf("http request failed: %w", res.Data.Error)
	}

	if err != nil {
		return chatroomResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	return res.Data, nil
}

func (gm *ChatroomManager) getChatroomDetailRequest(rids []string) (http.Request, error) {
	req := http.Request{
		URL:            gm.getChatroomDetailURL(rids),
		Method:         http.MethodGET,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + gm.client.appToken,
		},
	}
	return req, nil
}

func (gm *ChatroomManager) getChatroomDetailURL(rids []string) string {
	baseURL := gm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/chatrooms/%s", baseURL, strings.Join(rids, ","))
}
