/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ycj3/agora-chat-cli/log"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Manage push notifications",
	Long:  `Commands to manage push notifications.`,
}

var testPushCmd = &cobra.Command{
	Use:   "test",
	Short: "Test push notification",
	Example: heredoc.Doc(`
		# Send a test push notification for a specific user
		$ agchat push test --user <user-id>
	`),

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
		logger := log.NewLogger(verbose)
		handlePushTestResponse(res, logger)
		return nil
	},
}

func handlePushTestResponse(res ac.PushResponseResult, logger *log.Logger) {
	for _, dataItem := range res.Data {
		fields := map[string]interface{}{
			"timestamp": res.Timestamp,
		}

		if dataItem.Data != nil {
			jsonBytes, err := json.Marshal(dataItem.Data)
			if err != nil {
				logger.Error(fmt.Sprintf("Error marshalling map: %v", err), nil)
			}
			fields["data"] = fmt.Sprintf("the response body from FCM/APNS :%s", string(jsonBytes))
		}

		if dataItem.PushStatus == "SUCCESS" {
			logger.Info("Push notification success", fields)
		} else if dataItem.PushStatus == "FAIL" {
			fields["description"] = dataItem.Desc
			if dataItem.StatusCode > 0 {
				fields["statusCode"] = dataItem.StatusCode
			}
			logger.Error("Push notification failed", fields)
		}
	}
	toubleshootingMsg := "For more details, please refer to the response code documentation for the provider you are using: \n" +
		"FCM: https://firebase.google.com/docs/reference/fcm/rest/v1/ErrorCode\n" +
		"APNs :https://developer.apple.com/documentation/usernotifications/handling-notification-responses-from-apns"
	logger.Verbose(toubleshootingMsg, nil)
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.AddCommand(testPushCmd)

	testPushCmd.Flags().StringP("user", "u", "", "the user ID of the target user")
	testPushCmd.MarkFlagsRequiredTogether()
	testPushCmd.Flags().StringP("message", "m", `{"title": "Admin sent you a message", "content": "For push notification testing", "sub_title": "Test message is sent"}`, "JSON string for the push message")

	testPushCmd.MarkFlagRequired("user")
}
