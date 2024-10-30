/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/ycj3/agora-chat-cli/util"
)

type userDetailOptions struct {
	uID string
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
}

var onlineStatusCmd = &cobra.Command{
	Use:   "online-status",
	Short: "Query the online status of users",
	Long:  `Query the online status of users by specifying the users`,
	RunE: func(cmd *cobra.Command, args []string) error {

		users, _ := cmd.Flags().GetString("users")

		if users[len(users)-1] == ',' {
			logger.Warn("Extra spaces detected in --users flag. They will be removed.", nil)
			users = users[:len(users)-1]
		}

		userIDs := strings.Split(users, ",")

		statuses, err := client.User().UserOnlineStatuses(userIDs)

		if err != nil {
			return err
		}
		util.OutputJson(statuses)
		return nil
	},
}

func userDetailCmd() *cobra.Command {
	opts := &userDetailOptions{}

	var cmd = &cobra.Command{
		Use:   "detail",
		Short: "Get the detailed information of a specific user",
		Example: heredoc.Doc(`
				$ agchat user detail --user <user-id>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := client.User().QueryUser(opts.uID)
			if err != nil {
				return err
			}
			util.OutputJson(user)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.uID, "user", "u", "", "ID of the user")

	return cmd
}

func init() {

	userCmd.AddCommand(onlineStatusCmd)
	userCmd.AddCommand(userDetailCmd())

	onlineStatusCmd.Flags().StringP("users", "u", "", "Comma-separated list of users to query the online status for")
	onlineStatusCmd.MarkFlagRequired("users")

}
