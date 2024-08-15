/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/CarlsonYuan/agora-chat-cli/util"
	"github.com/spf13/cobra"
)

var provideCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage push providers added to an application",
}
var listProviderCmd = &cobra.Command{
	Use:   "list",
	Short: "list all providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := ac.NewClient()
		res, err := client.Provider().ListPushProviders()
		if err != nil {
			return fmt.Errorf("failed to list providers: %w", err)
		}

		if len(res.Entities) == 0 {
			fmt.Println("no provider added")
			return nil
		}
		util.OutputJson(res.Entities)
		return nil
	},
}

func init() {
	pushCmd.AddCommand(provideCmd)
	provideCmd.AddCommand(listProviderCmd)
}
