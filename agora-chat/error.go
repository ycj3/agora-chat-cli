/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import "fmt"

type Error struct {
	Exception        string `json:"exception,omitempty"`
	ErrorType        string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Exception: %s, ErrorType: %s, ErrorDescription: %s", e.Exception, e.ErrorType, e.ErrorDescription)
}
