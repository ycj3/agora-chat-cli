/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PushManager struct {
	client *Client
}

func (pm *PushManager) SendATestMessage(userID string, message map[string]interface{}) error {

	url := pm.client.appConfig.BaseURL + "/push/sync/" + userID

	pushRequest := map[string]interface{}{
		"strategy":    2,
		"pushMessage": message,
	}

	reqBody, err := json.Marshal(pushRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+pm.client.appToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting response: %v", err)
	}

	fmt.Printf("Response:\n%s\n", prettyJSON)

	return nil
}
