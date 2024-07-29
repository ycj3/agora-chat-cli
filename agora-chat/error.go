/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

type Error struct {
	Exception        string `json:"exception"`
	Duration         int    `json:"duration"`
	ErrorType        string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Timestamp        int64  `json:"timestamp"`
}
