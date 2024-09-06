/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
	"github.com/ycj3/agora-chat-cli/util"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage various attributes and actions of a user",
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

		client, err := ac.NewClient()
		if err != nil {
			logger.Error("Failed to get client", map[string]interface{}{
				"error": err.Error(),
				"desc":  "Please make sure you have created an app using the 'agchat apps --create' command",
			})
			return nil
		}
		statuses, err := client.User().UserOnlineStatuses(userIDs)

		if err != nil {
			return err
		}
		util.OutputJson(statuses)
		return nil
	},
}

func init() {

	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(onlineStatusCmd)

	onlineStatusCmd.Flags().StringP("users", "u", "", "Comma-separated list of users to query the online status for")
	onlineStatusCmd.MarkFlagRequired("users")

}
