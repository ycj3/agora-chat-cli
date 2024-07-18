package cmd

import (
	apps "github.com/CarlsonYuan/agora-chat-cli/apps"
	"github.com/spf13/cobra"
)

func appsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Manage all chat apps",
		RunE: func(cmd *cobra.Command, args []string) error {
			apps := apps.GetApps(cmd)

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
	cmd.Flags().BoolP("create", "c", false, "create a new chat application")
	cmd.Flags().BoolP("remove", "r", false, "Remove one or more application")
	cmd.Flags().BoolP("use", "", false, "set an active application for your working directory")

	return cmd
}
