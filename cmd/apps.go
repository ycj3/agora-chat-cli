package cmd

import (
	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
)

func appsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Manage all chat apps",
		RunE: func(cmd *cobra.Command, args []string) error {
			apps, _ := ac.LoadConfig()

			flist, _ := cmd.Flags().GetBool("list")
			if flist {
				return apps.ListAllApps()
			}

			fcreate, _ := cmd.Flags().GetBool("create")
			if fcreate {
				return apps.RunQuestionnaire()
			}

			fremove, _ := cmd.Flags().GetBool("remove")
			if fremove {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
				for _, appID := range args {
					if err := apps.Remove(appID); err != nil {
						return err
					}
					cmd.Println("Application successfully removed.")
				}
			}

			fuse, _ := cmd.Flags().GetBool("use")
			if fuse {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
				apps.Use(args[0])

			}
			return nil
		},
	}

	cmd.Flags().BoolP("list", "l", false, "list all chat applications")
	cmd.Flags().BoolP("create", "c", false, "register an existing chat application using details from the Agora console")
	cmd.Flags().BoolP("remove", "r", false, "remove one or more application")
	cmd.Flags().BoolP("use", "", false, "set an active application for your working directory")

	return cmd
}

func init() {
	rootCmd.AddCommand(appsCmd())
}
