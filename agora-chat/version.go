package agora_chat

import (
	"fmt"
)

const (
	versionMajor = 0
	versionMinor = 1
	versionPatch = 0
)

func FmtVersion() string {
	return fmt.Sprintf("%d.%d.%d",
		versionMajor,
		versionMinor,
		versionPatch)
}
