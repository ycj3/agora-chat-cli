/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate markdown documentation for the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("docs", 0755)
		if err != nil {
			fmt.Println("Error creating docs directory:", err)
			return
		}

		rootCmd.DisableAutoGenTag = true
		err = doc.GenMarkdownTree(rootCmd, "docs")
		if err != nil {
			fmt.Println("Error generating documentation:", err)
		} else {
			fmt.Println("Documentation generated in docs/ directory")
		}
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}
