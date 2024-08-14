/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"

	"github.com/CarlsonYuan/agora-chat-cli/http"
)

type DeviceManager struct {
	client *Client
}

type DeviceInfo struct {
	DeviceID     string `json:"device_id"`
	DeviceToken  string `json:"device_token"`
	NotifierName string `json:"notifier_name"`
}

type deviceResponseResult struct {
	Response
	Entities []DeviceInfo `json:"entities"`
}

func (dm *DeviceManager) AddPushDevice(userID, deviceID, deviceToken, notifierName string) ([]DeviceInfo, error) {
	req := dm.addPushDeviceRequest(userID, deviceID, deviceToken, notifierName)
	res, err := dm.client.deviceClient.Send(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Entities, nil
}

func (dm *DeviceManager) addPushDeviceRequest(userID, deviceID, deviceToken, notifierName string) http.Request {

	return http.Request{
		URL:            dm.pushBindURL(userID),
		Method:         http.MethodPUT,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + dm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: map[string]interface{}{
				"device_id":     deviceID,
				"device_token":  deviceToken,
				"notifier_name": notifierName,
			},
		},
	}
}

func (dm *DeviceManager) RemovePushDevice(userID, deviceID, notifierName string) ([]DeviceInfo, error) {
	req := dm.removePushDeviceRequest(userID, deviceID, notifierName)
	res, err := dm.client.deviceClient.Send(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Entities, nil
}

func (dm *DeviceManager) removePushDeviceRequest(userID, deviceID, notifierName string) http.Request {

	return http.Request{
		URL:            dm.pushBindURL(userID),
		Method:         http.MethodPUT,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + dm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: map[string]interface{}{
				"device_id":     deviceID,
				"device_token":  "",
				"notifier_name": notifierName,
			},
		},
	}
}

func (dm *DeviceManager) ListPushDevice(userID string) ([]DeviceInfo, error) {
	req := dm.listPushDeviceRequest(userID)
	res, err := dm.client.deviceClient.Send(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	return res.Data.Entities, nil
}

func (dm *DeviceManager) listPushDeviceRequest(userID string) http.Request {

	return http.Request{
		URL:            dm.pushBindURL(userID),
		Method:         http.MethodGET,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + dm.client.appToken,
		},
	}
}

func (dm *DeviceManager) pushBindURL(userID string) string {
	baseURL := dm.client.appConfig.BaseURL
	return fmt.Sprintf(baseURL+"/users/%s/push/binding", userID)
}
