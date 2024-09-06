/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

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
	Short: "Test whether the push notification credentials and notification services work properly",
	Example: heredoc.Doc(`
		# Send a test push notification to a specific user
		$ agchat push test --user <user-id>
	`),

	RunE: func(cmd *cobra.Command, args []string) error {

		userID, _ := cmd.Flags().GetString("user")
		if userID == "" {
			cmd.Println("Usage: agchat push test --user <user-id>")
		}

		message, _ := cmd.Flags().GetString("message")
		var msg ac.PushMessage
		if err := json.Unmarshal([]byte(message), &msg); err != nil {
			logger.Error("failed to parse push message JSON", map[string]interface{}{
				"error": err,
			})
		}

		client, err := ac.NewClient()
		if err != nil {
			logger.Error("Failed to get client", map[string]interface{}{
				"error": err.Error(),
				"desc":  "Please make sure you have created an app using the 'agchat apps --create' command",
			})
			return nil
		}

		// Step 1: check if you have registered the push notification credentials with the Agora Chat Server.
		err = checkPushCredential(client, cmd, args)
		if err != nil {
			logger.Error("✖ Step 1: Failed to check the push notification credentials %w", map[string]interface{}{
				"error":   err.Error(),
				"success": false,
			})
			return nil
		}

		// Step 2: check if you have registered a device token with the Agora Chat Server for the target user.
		err = checkPushDeviceToken(client, userID, cmd, args)
		if err != nil {
			logger.Error("✖ Step 2: Failed to check the push device tokens", map[string]interface{}{
				"error":   err.Error(),
				"success": false,
			})
			return nil
		}

		// Step 3: Send push notification
		res, err := client.Push().SyncPush(userID, ac.OnlyPushPrivider, msg)

		err = handlePushTestResponse(res)
		if err != nil {
			return nil
		}
		logger.Info("Push notification test completed.", nil)
		return nil

	},
}

// Step 1
func checkPushCredential(c ac.Client, cmd *cobra.Command, args []string) error {
	res, err := c.Provider().ListPushProviders()
	if err != nil {
		return fmt.Errorf("Failed to list providers: %w", err)
	}
	if len(res.Entities) == 0 {
		return fmt.Errorf("No push notification credentials registered with the Agora Chat Server.")
	}
	logger.Info("✔ Step 1: Checked the push notification credentials registered with the Agora Chat Server", map[string]interface{}{
		"count": len(res.Entities),
	})
	return nil
}

// Step 2
func checkPushDeviceToken(c ac.Client, userID string, cmd *cobra.Command, args []string) error {
	devices, err := c.Device().ListPushDevice(userID)
	if err != nil {
		return fmt.Errorf("Failed to list push devices: %w", err)
	}
	if len(devices) == 0 {
		return fmt.Errorf("No Devices registered with the Agora Chat Server.")
	}
	logger.Info("✔ Step 2: Checked the push device tokens registered with the Agora Chat Server", map[string]interface{}{
		"count": len(devices),
	})
	return nil
}

// Step 3
func handlePushTestResponse(res ac.PushResponseResult) error {

	// First, handle all success entries
	var succssEntries []ac.PushResult
	for _, dataItem := range res.Data {
		if dataItem.PushStatus == "SUCCESS" {
			succssEntries = append(succssEntries, dataItem)
		}
	}

	// Then, handle all failure entries
	var failuresEntries []ac.PushResult
	for _, dataItem := range res.Data {
		if dataItem.PushStatus == "FAIL" {
			failuresEntries = append(failuresEntries, dataItem)
		}
	}

	if len(succssEntries) > 0 {
		logger.Info("✔ Step 3: Sent push notification\n", map[string]interface{}{
			"success": map[string]interface{}{"results": succssEntries, "count": len(succssEntries)},
			"failure": map[string]interface{}{"errorResults": failuresEntries, "count": len(failuresEntries)},
			"count":   len(res.Data),
		})
	} else {
		logger.Error("✖ Step 3: Failed to send push notification\n", map[string]interface{}{
			"failure": map[string]interface{}{"errorResults": failuresEntries, "count": len(failuresEntries)},
			"count":   len(res.Data),
		})
		return fmt.Errorf("")
	}

	toubleshootingMsg := "For more details, please refer to the response code documentation for the provider you are using: \n" +
		"FCM: https://firebase.google.com/docs/reference/fcm/rest/v1/ErrorCode\n" +
		"APNs :https://developer.apple.com/documentation/usernotifications/handling-notification-responses-from-apns"
	logger.Verbose(toubleshootingMsg, nil)
	return nil
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.AddCommand(testPushCmd)

	testPushCmd.Flags().StringP("user", "u", "", "the user ID of the target user")
	testPushCmd.MarkFlagsRequiredTogether()
	testPushCmd.Flags().StringP("message", "m", `{"title": "Admin sent you a message", "content": "For push notification testing", "sub_title": "Test message is sent"}`, "JSON string for the push message")

	testPushCmd.MarkFlagRequired("user")
}
