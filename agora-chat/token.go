/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/chatTokenBuilder"
)

type TokenManager struct {
	client *Client
}

func NewTokenManager(client *Client) *TokenManager {
	return &TokenManager{client: client}
}

func (tm *TokenManager) generateChatAppToken() (string, error) {
	return chatTokenBuilder.BuildChatAppToken(tm.client.appConfig.AppID, tm.client.appConfig.AppCertificate, tm.client.appTokenExp)
}

func (tm *TokenManager) GenerateChatUserToken(userID string) (string, error) {
	return chatTokenBuilder.BuildChatUserToken(tm.client.appConfig.AppID, tm.client.appConfig.AppCertificate, userID, tm.client.appTokenExp)
}
