/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

type MessageType string

const (
	MessageTypeText   MessageType = "txt"    // Text message
	MessageTypeImage  MessageType = "img"    // Image message
	MessageTypeAudio  MessageType = "audio"  // Voice message
	MessageTypeVideo  MessageType = "video"  // Video message
	MessageTypeFile   MessageType = "file"   // File message
	MessageTypeLoc    MessageType = "loc"    // Location message
	MessageTypeCmd    MessageType = "cmd"    // Command message
	MessageTypeCustom MessageType = "custom" // Custom message
)

// TextMessageBody represents the payload data for a text type message.
type TextMessageBody struct {
	Msg string `json:"msg,omitempty"`
}

// CMDMessageBody represents the payload data of the command type message.
type CMDMessageBody struct {
	Action string `json:"action,omitempty"`
}

// LocationMessageBody represents the payload data of the location type message.
type LocationMessageBody struct {
	Lat  string `json:"lat,omitempty"`
	Lng  string `json:"lng,omitempty"`
	Addr string `json:"addr,omitempty"`
}

// CustomMessageBody represents the payload data of the custom type message.
type CustomMessageBody struct {
	CustomEvent string            `json:"customEvent,omitempty"`
	CustomExts  map[string]string `json:"customExts,omitempty"`
}

type Message struct {
	From string      `json:"from,omitempty"`
	To   []string    `json:"to,omitempty"`
	Type MessageType `json:"type,omitempty"`
	Body interface{} `json:"body,omitempty"`
}
