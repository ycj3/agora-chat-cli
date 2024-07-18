/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/CarlsonYuan/agora-chat-cli/apps"
	"github.com/CarlsonYuan/agora-chat-cli/version"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var cfgPath = new(string)

func rootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "agorachat <command> <subcommand> [flags]",
		Short: "Agora Chat CLI",
		Long:  "Interact with your Agora Chat applications easily",
		Example: heredoc.Doc(`
	
		`),
		Version: version.FmtVersion(),
	}

	root.AddCommand(
		appsCmd(),
	)

	cobra.OnInitialize(apps.GetInitConfig(root, cfgPath))

	root.SetOut(os.Stdout)

	return root
}

func Execute() {
	err := rootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
