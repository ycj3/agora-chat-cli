/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
	"github.com/ycj3/agora-chat-cli/log"
)

var (
	verbose bool
)

var version = "dev"
var logger *log.Logger

var rootCmd = &cobra.Command{
	Use:   "agchat <command> <subcommand> [flags]",
	Short: "Agora Chat CLI",
	Long:  "Interact with your Agora Chat applications easily",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = log.NewLogger(verbose)
	},
	Example: heredoc.Doc(`

	`),
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	if _, err := ac.LoadConfig(); err != nil {
		rootCmd.PrintErrf("Error loading config: %v\n", err)
	}
}
