/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	"strings"

	"github.com/ycj3/agora-chat-cli/http"
)

type affiliation struct {
	Member string `json:"member,omitempty"`
	Owner  string `json:"owner,omitempty"`
}

type groupDetail struct {
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

type groupResponseResult struct {
	Response
	Path            string        `json:"path"`
	URI             string        `json:"uri"`
	Organization    string        `json:"organization"`
	Application     string        `json:"application"`
	Data            []groupDetail `json:"data"`
	ApplicationName string        `json:"applicationName"`
}

type GroupManager struct {
	client *client
}

func (gm *GroupManager) GetGroupDetail(gids []string) (groupResponseResult, error) {
	request, err := gm.getGroupDetailRequest(gids)
	if err != nil {
		return groupResponseResult{}, err
	}

	res, err := gm.client.groupClient.Send(request)
	if err != nil {
		return groupResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	return res.Data, nil
}

func (gm *GroupManager) getGroupDetailRequest(gids []string) (http.Request, error) {
	req := http.Request{
		URL:            gm.getGroupDetailURL(gids),
		Method:         http.MethodGET,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + gm.client.appToken,
		},
	}
	return req, nil
}

func (gm *GroupManager) getGroupDetailURL(gids []string) string {
	baseURL := gm.client.appConfig.BaseURL
	return fmt.Sprintf("%s/chatgroups/%s", baseURL, strings.Join(gids, ","))
}
