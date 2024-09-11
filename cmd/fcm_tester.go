/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

var fcmCmd = &cobra.Command{
	Use:   "fcm",
	Short: "Test FCM push notifications",
}

var testFCMPushCmd = &cobra.Command{
	Use:   "test",
	Short: "Test whether the push notification credentials and notification services work properly in FCM",
	Example: heredoc.Doc(`
		# Send a test message to a specific device
		$ agchat fcm test --key <service-account-key> --token <device-token>
	`),

	Run: func(cmd *cobra.Command, args []string) {

		kp, _ := cmd.Flags().GetString("key")
		rt, _ := cmd.Flags().GetString("token")
		msg, _ := cmd.Flags().GetString("message")

		app := InitializeFirebase(kp)

		// Obtain a messaging.Client from the App.
		ctx := context.Background()
		client, err := app.Messaging(ctx)
		if err != nil {
			logger.Fatal("Failed to get Messaging client", map[string]interface{}{
				"error": err,
			})
		}

		// This registration token comes from the client FCM SDKs.
		registrationToken := rt

		var noti messaging.Notification
		if err := json.Unmarshal([]byte(msg), &noti); err != nil {
			logger.Error("Failed to parse message ", map[string]interface{}{
				"error": err,
			})
			return
		}
		message := &messaging.Message{
			Notification: &noti,
			Token:        registrationToken,
		}

		// Send a message to the device corresponding to the provided
		// registration token.
		response, err := client.Send(ctx, message)
		if err != nil {
			logger.Error("Failed to send message", map[string]interface{}{
				"error":       err.Error(),
				"description": "description and resolution steps, see https://firebase.google.com/docs/reference/fcm/rest/v1/ErrorCode",
			})
			return
		}
		// Response is a message ID string.
		logger.Info("Successfully sent message", map[string]interface{}{
			"message-id-string": response,
		})
	},
}

func InitializeFirebase(keyPath string) *firebase.App {
	opt := option.WithCredentialsFile(keyPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error("Failed to initialize app", map[string]interface{}{
			"error": err,
		})
	}
	return app
}

func init() {
	fcmCmd.AddCommand(testFCMPushCmd)

	testFCMPushCmd.Flags().StringP("token", "t", "", "The device's registration token")
	testFCMPushCmd.Flags().StringP("key", "k", "", "The service account JSON file")
	testFCMPushCmd.Flags().StringP("message", "m", `{"title": "FCM Message", "body": "This is an FCM notification message!"}`, "JSON string for the push message")

	testFCMPushCmd.MarkFlagRequired("token")
	testFCMPushCmd.MarkFlagRequired("key")
}
