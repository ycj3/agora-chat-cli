/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/CarlsonYuan/agora-chat-cli/config"
	"github.com/CarlsonYuan/agora-chat-cli/version"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agorachat <command> <subcommand> [flags]",
	Short: "Agora Chat CLI",
	Long:  "Interact with your Agora Chat applications easily",
	Example: heredoc.Doc(`

	`),
	Version: version.FmtVersion(),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize()

	if _, err := config.LoadConfig(); err != nil {
		rootCmd.PrintErrf("Error loading config: %v\n", err)
	}
}
