/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	gohttp "net/http"

	"github.com/CarlsonYuan/agora-chat-cli/http"
)

type UserManager struct {
	client *Client
}

type userOnlineStatus map[string]string

type userResponseResult struct {
	Error
	response
	Data []userOnlineStatus `json:"data"`
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
