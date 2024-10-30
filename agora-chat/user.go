/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	gohttp "net/http"

	"github.com/ycj3/agora-chat-cli/http"
)

type User struct {
	Created  int64 `json:"created"`
	Modified int64 `json:"modified"`

	Activated bool `json:"activated"`

	ID   string `json:"username"`
	UUID string `json:"uuid"`

	NotificationDisplayStyle *int       `json:"notification_display_style,omitempty"`
	PushInfo                 []PushInfo `json:"pushInfo,omitempty"`
}

type PushInfo struct {
	DeviceID     string `json:"device_Id"`
	DeviceToken  string `json:"device_token"`
	NotifierName string `json:"notifier_name"`
}

type UserManager struct {
	client *client
}

type userOnlineStatus map[string]string

type userResponseResult struct {
	Error
	Response
	Data  []userOnlineStatus `json:"data,omitempty"`
	Users []User             `json:"entities,omitempty"`
}

func (um *UserManager) UserOnlineStatuses(userIDs []string) ([]userOnlineStatus, error) {
	request := um.userOnlineStatusesRequest(userIDs)
	res, err := um.client.userClient.Send(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	if res.StatusCode != gohttp.StatusOK {
		return nil, res.Data.Error
	}
	return res.Data.Data, nil
}

func (um *UserManager) userOnlineStatusesRequest(userIDs []string) http.Request {

	return http.Request{
		URL:            um.userOnlineStatusesURL(),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + um.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: map[string]interface{}{
				"usernames": userIDs,
			},
		},
	}
}

func (um *UserManager) userOnlineStatusesURL() string {
	baseURL := um.client.appConfig.BaseURL
	return fmt.Sprintf(baseURL + "/users/batch/status")
}

func (um *UserManager) QueryUser(uID string) (User, error) {
	request := um.queryUserRequest(uID)
	res, err := um.client.userClient.Send(request)
	if err != nil {
		return User{}, fmt.Errorf("request failed: %w", err)
	}
	if res.StatusCode != gohttp.StatusOK {
		return User{}, res.Data.Error
	}
	return res.Data.Users[0], nil
}

func (um *UserManager) queryUserRequest(uID string) http.Request {
	return http.Request{
		URL:            um.queryUserURL(uID),
		Method:         http.MethodGET,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + um.client.appToken,
		},
	}
}

func (um *UserManager) queryUserURL(uID string) string {
	baseURL := um.client.appConfig.BaseURL
	return fmt.Sprintf(baseURL+"/users/%s", uID)
}
