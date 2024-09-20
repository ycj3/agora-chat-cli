package cmdutil

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func Test_IsAuthCheckEnabled(t *testing.T) {
	tests := []struct {
		name               string
		init               func() (*cobra.Command, error)
		isAuthCheckEnabled bool
	}{
		{
			name: "no annotations",
			init: func() (*cobra.Command, error) {
				cmd := &cobra.Command{}
				cmd.Flags().Bool("flag", false, "")
				return cmd, nil
			},
			isAuthCheckEnabled: true,
		},
		{
			name: "command-level disable",
			init: func() (*cobra.Command, error) {
				cmd := &cobra.Command{}
				DisableAuthCheck(cmd)
				return cmd, nil
			},
			isAuthCheckEnabled: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := tt.init()
			require.NoError(t, err)

			// IsAuthCheckEnabled assumes commands under test are subcommands
			parent := &cobra.Command{Use: "root"}
			parent.AddCommand(cmd)
			require.Equal(t, tt.isAuthCheckEnabled, IsAuthCheckEnabled(cmd))
		})
	}
}
