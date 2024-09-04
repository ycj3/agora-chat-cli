package agora_chat

import (
	"fmt"
)

const (
	versionMajor = 0
	versionMinor = 2
	versionPatch = 1
)

func FmtVersion() string {
	return fmt.Sprintf("%d.%d.%d",
		versionMajor,
		versionMinor,
		versionPatch)
}
