/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Manage Chat SDK logs of a user",
}

var logUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Send a upload command to the online users",
	Long:  `Send a upload command to the online users by specifying the users`,
	RunE: func(cmd *cobra.Command, args []string) error {

		users, _ := cmd.Flags().GetString("users")

		if users[len(users)-1] == ',' {
			fmt.Println("Warning: Extra spaces detected in --users flag. They will be removed.")
			users = users[:len(users)-1]
		}

		userIDs := strings.Split(users, ",")

		statuses, err := client.User().UserOnlineStatuses(userIDs)

		if err != nil {
			return err
		}

		for _, status := range statuses {
			for user, onlineStatus := range status {
				fmt.Printf("User: %s, Online Status: %s\n", user, onlineStatus)
				if onlineStatus == "online" {
					sendUploadLogsCMD([]string{user})
				}
			}
		}

		return nil
	},
}

func sendUploadLogsCMD(userIDs []string) {
	message := ac.Message{
		From: "admin",
		To:   userIDs,
		Type: ac.MessageTypeCmd,
		Body: ac.CMDMessageBody{
			Action: "em_upload_log",
		},
	}
	msgIDs, err := client.Message().SendUsersMessage(&message)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range msgIDs {
		fmt.Printf("upload command sent successfully to %s, message id is: %s\n", key, value)
	}
}
func init() {

	logCmd.AddCommand(logUploadCmd)

	logUploadCmd.Flags().String("users", "", "Comma-separated list of users to send a upload command to the online users for")
	logUploadCmd.MarkFlagRequired("users")

}
