package cmdutil

import "github.com/spf13/cobra"

const skipAuthCheckAnnotation = "skipAuthCheck"

func DisableAuthCheck(cmd *cobra.Command) {
	if cmd.Annotations == nil {
		cmd.Annotations = map[string]string{}
	}

	cmd.Annotations[skipAuthCheckAnnotation] = "true"
}

func IsAuthCheckEnabled(cmd *cobra.Command) bool {
	for c := cmd; c.Parent() != nil; c = c.Parent() {
		// Check whether any command marked as DisableAuthCheck is set
		if c.Annotations != nil && c.Annotations[skipAuthCheckAnnotation] == "true" {
			return false
		}
	}

	return true
}
