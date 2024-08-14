/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

type Response struct {
	Timestamp int64  `json:"timestamp"`
	Duration  int    `json:"duration"`
	Action    string `json:"action,omitempty"`
}
