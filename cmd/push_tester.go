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

		// Step 1: check if you have registered the push notification credentials with the Agora Chat Server.
		err := checkPushCredential()
		if err != nil {
			logger.Error("Failed to check push notification credentials", map[string]interface{}{
				"error": err.Error(),
			})
			return nil
		}

		// Step 2: check if you have registered a device token with the Agora Chat Server for the target user.
		err = checkPushDeviceToken(userID)
		if err != nil {
			logger.Error("Failed to check device token(s)", map[string]interface{}{
				"error": err.Error(),
			})
			return nil
		}

		// Step 3: Send push notification
		res, _ := client.Push().SyncPush(userID, ac.OnlyPushPrivider, msg)

		err = handlePushTestResponse(res)
		if err != nil {
			return err
		}
		return nil

	},
}

// Step 1
func checkPushCredential() error {
	res, err := client.Provider().ListPushProviders()
	if err != nil {
		return fmt.Errorf("failed to list providers: %w", err)
	}
	if len(res.Entities) == 0 {
		return fmt.Errorf("no push notification credentials found in the current active app")
	}
	return nil
}

// Step 2
func checkPushDeviceToken(userID string) error {
	devices, err := client.Device().ListPushDevice(userID)
	if err != nil {
		return fmt.Errorf("filed to list push devices: %w", err)
	}

	if len(devices) == 0 {
		return fmt.Errorf("no device tokens registered for uid: %s", userID)
	}

	if len(devices) > 1 {
		logger.Info(fmt.Sprintf("%d device token(s) registered", len(devices)), map[string]interface{}{
			"uid": userID,
		})
	}
	return nil
}

// Step 3

func handlePushTestResponse(res ac.PushResponseResult) error {

	// var succssEntries []ac.PushResult
	var dot string
	if len(res.Data) > 1 {
		dot = " - "
	} else {
		dot = ""
	}

	for _, dataItem := range res.Data {
		if dataItem.PushStatus == "SUCCESS" {
			if dataItem.Data.Name != "" {
				logger.Info(fmt.Sprintf("%sMessage sent successfully via FCM", dot), map[string]interface{}{
					"message-id": dataItem.Data.Name,
				})
			} else if dataItem.Data.StatusCode == 200 && dataItem.Data.ApnsUniqueId != "" && dataItem.Data.Accepted {
				logger.Info(fmt.Sprintf("%sMessage sent successfully via APNs", dot), map[string]interface{}{
					"apnsUniqueId": dataItem.Data.ApnsUniqueId,
				})
			} else {
				return fmt.Errorf("unknow success type for push")
			}
		}
	}

	for _, dataItem := range res.Data {
		if dataItem.PushStatus == "FAIL" {
			if dataItem.Data.FcmError != nil {
				logger.Error(fmt.Sprintf("%sFailed to send message via FCM: %s", dot, dataItem.Data.FcmError.Message), map[string]interface{}{
					"code":      dataItem.Data.FcmError.Code,
					"errorCode": dataItem.Data.FcmError.Details[0].ErrorCode,
					"status":    dataItem.Data.FcmError.Status,
				})
			} else if dataItem.Data.StatusCode != 200 && dataItem.Data.ApnsUniqueId == "" && !dataItem.Data.Accepted {
				logger.Error(fmt.Sprintf("%sFailed to send message via APNs: %s", dot, dataItem.Data.RejectionReason), map[string]interface{}{
					"statusCode": dataItem.Data.StatusCode,
					"apnsId":     dataItem.Data.ApnsId,
					"token":      dataItem.Data.PushNotification.Token,
				})
			} else if dataItem.Desc != "" {
				logger.Error(fmt.Sprintf("%sFailed to send message: %s", dot, dataItem.Desc), nil)
			} else {
				return fmt.Errorf("unknow failure type for push")
			}

		}
	}

	toubleshootingMsg := "For more details, please refer to the response code documentation for the provider you are using: \n" +
		"FCM: https://firebase.google.com/docs/reference/fcm/rest/v1/ErrorCode\n" +
		"APNs :https://developer.apple.com/documentation/usernotifications/handling-notification-responses-from-apns"
	logger.Verbose(toubleshootingMsg, nil)
	return nil
}

func init() {
	pushCmd.AddCommand(testPushCmd)

	testPushCmd.Flags().StringP("user", "u", "", "the user ID of the target user")
	testPushCmd.MarkFlagsRequiredTogether()
	testPushCmd.Flags().StringP("message", "m", `{"title": "Admin sent you a message", "content": "For push notification testing", "sub_title": "Test message is sent"}`, "JSON string for the push message")

	testPushCmd.MarkFlagRequired("user")
}
