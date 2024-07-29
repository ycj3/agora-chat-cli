/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Manage device information bound to a user",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all device info",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetString("user")

		client := ac.GetActiveApp().GetClient()
		client.Device().List(userID)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)
	deviceCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("user", "u", "", "the user ID of the target user")
	listCmd.MarkFlagRequired("user")

}
