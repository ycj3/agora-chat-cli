/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Manage push notifications",
	Long:  `Commands to manage push notifications.`,
}

var testPushCmd = &cobra.Command{
	Use:   "test",
	Short: "Test push notification",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetString("user")
		message, _ := cmd.Flags().GetString("message")

		var pushMessage map[string]interface{}

		if err := json.Unmarshal([]byte(message), &pushMessage); err != nil {
			return fmt.Errorf("failed to parse push message JSON: %v", err)
		}

		if userID == "" {
			cmd.Println("Usage: agchat push test --user <user-id>")
			return nil
		}

		apps, err := ac.LoadConfig()
		if err != nil {
			return err
		}

		active, err := apps.GetActiveApp()
		if err != nil {
			return err
		}

		client, err := ac.NewClient(active)
		if err != nil {
			return err
		}
		client.Push().SendATestMessage(userID, pushMessage)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.AddCommand(testPushCmd)

	testPushCmd.Flags().StringP("user", "u", "", "the user ID of the target user")
	testPushCmd.MarkFlagsRequiredTogether()
	testPushCmd.Flags().StringP("message", "m", `{"title": "Admin sent you a message", "content": "For push notification testing", "sub_title": "Test message is sent"}`, "JSON string for the push message")

	testPushCmd.MarkFlagRequired("user")
}
