/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
)

func validateMessage(message *Message) error {
	if message == nil {
		return fmt.Errorf("message must not be nil")
	}

	if message.To == nil || len(message.To) <= 0 {
		return fmt.Errorf("to must be specified")
	}

	if message.Type == "" {
		return fmt.Errorf("type must be specified")
	}

	if message.Body == nil {
		return fmt.Errorf("body must be specified")
	}

	if message.Type == MessageTypeText {
		textBody, ok := message.Body.(TextMessageBody)
		if !ok {
			return fmt.Errorf("invalid type: expected TextMessageBody")
		}
		return validateTextMessageBody(textBody)
	}

	if message.Type == MessageTypeCmd {
		cmdBody, ok := message.Body.(CMDMessageBody)
		if !ok {
			return fmt.Errorf("invalid type: expected CMDMessageBody")
		}
		return validateCMDMessageBody(cmdBody)
	}

	if message.Type == MessageTypeLoc {
		locBody, ok := message.Body.(LocationMessageBody)
		if !ok {
			return fmt.Errorf("invalid type: expected CMDMessageBody")
		}
		return validateLocationMessageBody(locBody)
	}

	return nil

}

func validateTextMessageBody(body TextMessageBody) error {
	if body.Msg == "" {
		return fmt.Errorf("msg is required when specifying TextMessageBody")
	}
	return nil
}

func validateCMDMessageBody(body CMDMessageBody) error {
	if body.Action == "" {
		return fmt.Errorf("action is required when specifying CMDMessageBody")
	}
	return nil
}

func validateLocationMessageBody(body LocationMessageBody) error {
	if body.Lat == "" {
		return fmt.Errorf("lat is required when specifying LocationMessageBody")
	}

	if body.Lng == "" {
		return fmt.Errorf("lng is required when specifying LocationMessageBody")
	}

	if body.Addr == "" {
		return fmt.Errorf("addr is required when specifying LocationMessageBody")
	}

	return nil
}
