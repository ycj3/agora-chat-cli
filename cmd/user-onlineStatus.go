/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/spf13/cobra"
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
			fmt.Println("Warning: Extra spaces detected in --users flag. They will be removed.")
			users = users[:len(users)-1]
		}

		userIDs := strings.Split(users, ",")

		client := ac.NewClient()
		statuses, err := client.User().UserOnlineStatuses(userIDs)

		if err != nil {
			return err
		}

		for _, status := range statuses {
			for user, onlineStatus := range status {
				fmt.Printf("User: %s, Online Status: %s\n", user, onlineStatus)
			}
		}

		return nil
	},
}

func init() {

	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(onlineStatusCmd)

	onlineStatusCmd.Flags().String("users", "", "Comma-separated list of users to query the online status for")
	onlineStatusCmd.MarkFlagRequired("users")

}
