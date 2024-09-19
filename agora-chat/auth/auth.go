/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package builder

import (
	"fmt"
	"os"

	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/chatTokenBuilder"
)

const (
	acToken = "AC_TOKEN"
)

type Auth interface {
	TokenFromEnvOrGenerate() (string, error)
	AppTokenFromBuilder() (string, error)
	UserTokenFromBuilder(userID string) (string, error)
	HasEnvToken() bool
}

type builder struct {
	AppID          string
	AppCertificate string
	AppTokenExp    uint32
}

func NewAuth(appID, appCertificate string, appTokenExp uint32) (*builder, error) {
	//TODO check params

	return &builder{
		AppID:          appID,
		AppCertificate: appCertificate,
		AppTokenExp:    appTokenExp,
	}, nil
}

func (t *builder) TokenFromEnvOrBuilder() (string, error) {
	if token := TokenFromEnv(); token != "" {
		return token, nil
	}

	token, err := t.AppTokenFromBuilder()

	if err != nil {
		return "", fmt.Errorf("failed to generate app token :%s", err)
	}

	return token, nil
}

func TokenFromEnv() string {
	if token := os.Getenv(acToken); token != "" {
		return token
	}
	return ""
}

func HasEnvToken() bool {
	token := TokenFromEnv()
	return token != ""
}

func (t *builder) AppTokenFromBuilder() (string, error) {
	return chatTokenBuilder.BuildChatAppToken(t.AppID, t.AppCertificate, t.AppTokenExp)
}

func (t *builder) UserTokenFromBuilder(userID string) (string, error) {
	return chatTokenBuilder.BuildChatUserToken(t.AppID, t.AppCertificate, userID, t.AppTokenExp)
}
