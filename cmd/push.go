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
		if userID == "" {
			cmd.Println("Usage: agchat push test --user <user-id>")
			return nil
		}

		message, _ := cmd.Flags().GetString("message")
		var msg ac.PushMessage
		if err := json.Unmarshal([]byte(message), &msg); err != nil {
			return fmt.Errorf("failed to parse push message JSON: %v", err)
		}

		client := ac.NewClient()

		res, err := client.Push().SyncPush(userID, ac.OnlyPushPrivider, msg)
		if err != nil {
			return fmt.Errorf("failed to send push notification: %w", err)
		}

		for _, pushResult := range res {
			switch pushResult.PushStatus {
			case "SUCCESS":
				if pushResult.Data != nil {
					cmd.Println("Success - PushSuccessResult:")
					if pushResult.Data.Result != "" {
						cmd.Printf("Result: %s\n", pushResult.Data.Result)
					}
					if pushResult.Data.MsgID != nil {
						if len(*pushResult.Data.MsgID) > 0 {
							cmd.Printf("MsgID: %+v\n", *pushResult.Data.MsgID)
						} else {
							cmd.Println("MsgID is empty")
						}
					} else {
						cmd.Println("MsgID is nil")
					}
				} else {
					cmd.Println("Success - Data is nil")
				}
			case "FAIL":
				if pushResult.Desc != nil {
					cmd.Printf("Failure - Desc: %s\n", *pushResult.Desc)
				} else {
					cmd.Println("Failure - No description provided")
				}
			default:
				cmd.Println("Unknown pushStatus:", pushResult.PushStatus)
			}
		}
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
