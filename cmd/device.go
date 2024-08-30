/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
	"github.com/ycj3/agora-chat-cli/util"
)

var device ac.DeviceInfo

func init() {

	rootCmd.AddCommand(deviceCmd)
	deviceCmd.PersistentFlags().StringP("user", "u", "", "the user ID of the target user")
	deviceCmd.MarkPersistentFlagRequired("user")

	deviceCmd.AddCommand(addDeviceCmd)
	addDeviceCmd.Flags().StringVar(&device.DeviceID, "device-id", "", "Device ID")
	addDeviceCmd.Flags().StringVar(&device.NotifierName, "notifier-name", "", "Notifier Name")
	addDeviceCmd.Flags().StringVar(&device.DeviceToken, "device-token", "", "Device Token")
	addDeviceCmd.MarkFlagRequired("device-id")
	addDeviceCmd.MarkFlagRequired("notifier-name")
	addDeviceCmd.MarkFlagRequired("device-token")

	deviceCmd.AddCommand(removeDeviceCmd)
	removeDeviceCmd.Flags().StringVar(&device.DeviceID, "device-id", "", "Device ID")
	removeDeviceCmd.Flags().StringVar(&device.NotifierName, "notifier-name", "", "Notifier Name")
	removeDeviceCmd.MarkFlagRequired("device-id")
	removeDeviceCmd.MarkFlagRequired("notifier-name")

	deviceCmd.AddCommand(listDevicesCmd)
}

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Manage device information bound to a user",
}

var addDeviceCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new device",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetString("user")

		client := ac.NewClient()
		devices, err := client.Device().AddPushDevice(userID, device.DeviceID, device.DeviceToken, device.NotifierName)
		if err != nil {
			return fmt.Errorf("failed to add device: %w", err)
		}

		for i, device := range devices {
			cmd.Printf("Device %d: %+v\n", i+1, device)
		}

		return nil
	},
}

var removeDeviceCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an existing device",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetString("user")

		client := ac.NewClient()
		_, err := client.Device().RemovePushDevice(userID, device.DeviceID, device.NotifierName)
		if err != nil {
			return fmt.Errorf("failed to remove device: %w", err)
		}
		fmt.Printf("Removed Device with ID: %s\n", device.DeviceID)
		return nil
	},
}

var listDevicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all devices",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetString("user")

		client := ac.NewClient()
		devices, err := client.Device().ListPushDevice(userID)
		if err != nil {
			return fmt.Errorf("failed to list devices: %w", err)
		}

		if len(devices) == 0 {
			fmt.Println("no device registered")
			return nil
		}
		util.OutputJson(devices)
		return nil
	},
}
