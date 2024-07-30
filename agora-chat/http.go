/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

type response struct {
	Timestamp int64  `json:"timestamp"`
	Duration  int    `json:"duration"`
	Action    string `json:"action,omitempty"`
}
