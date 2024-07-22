/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agchat <command> <subcommand> [flags]",
	Short: "Agora Chat CLI",
	Long:  "Interact with your Agora Chat applications easily",
	Example: heredoc.Doc(`

	`),
	Version: ac.FmtVersion(),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize()

	if _, err := ac.LoadConfig(); err != nil {
		rootCmd.PrintErrf("Error loading config: %v\n", err)
	}
}
