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
var client ac.Client

func rootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "agchat <command> <subcommand> [flags]",
		Short: "Agora Chat CLI",
		Long:  "Interact with your Agora Chat applications easily",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger = log.NewLogger(verbose)
			if cmd.Use != "apps" && cmd.Use != "doc" {
				initChatClient()
			}

		},
		Example: heredoc.Doc(`

		`),
		Version: version,
	}

	cmd.DisableAutoGenTag = true

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	if _, err := ac.LoadConfig(); err != nil {
		cmd.PrintErrf("Error loading config: %v\n", err)
	}

	cmd.AddCommand(appsCmd())
	cmd.AddCommand(docCmd())

	cmd.AddCommand(deviceCmd)
	cmd.AddCommand(logCmd)
	cmd.AddCommand(pushCmd)
	cmd.AddCommand(userCmd)
	cmd.AddCommand(tokenCmd)

	return cmd
}

func Execute() int {
	err := rootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
	return 0
}

func init() {
	cobra.OnInitialize()
}

func initChatClient() {
	var err error
	client, err = ac.NewClient()
	if err != nil {
		logger.Fatal("Failed to get client", map[string]interface{}{
			"error": err.Error(),
			"desc":  "Please make sure you have created an app using the 'agchat apps --create' command",
		})
	}
}
