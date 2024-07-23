/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/accesstoken2"
	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate and parse agora tokens",
	RunE: func(cmd *cobra.Command, args []string) error {
		apps, err := ac.LoadConfig()
		if err != nil {
			return err
		}

		if token, _ := cmd.Flags().GetString("parse"); token != "" {
			tokenInfo, err := parseToken(token)
			if err != nil {
				return err
			}
			cmd.Printf("Token information:\n%s\n", tokenInfo)
		}

		active, err := apps.GetActiveApp()
		if err != nil {
			return err
		}

		client, err := ac.NewClient(active)
		if err != nil {
			return err
		}

		if userID, _ := cmd.Flags().GetString("user"); userID != "" {
			userToken, err := client.Tokens().GenerateChatUserToken(userID)
			if err != nil {
				return err
			}
			cmd.Printf("Token for user [%s]:\n%s\n", userID, userToken)
		}
		return nil
	},
}

func parseToken(token string) (string, error) {
	version := token[:3]
	if version != "007" {
		return "", fmt.Errorf("not supported, only token version 007 is supported")
	}

	token = token[3:]
	accesstoken := accesstoken2.CreateAccessToken()
	finalToken := version + cleanBase64(token)
	result, err := accesstoken.Parse(finalToken)
	if !result {
		return "", err
	}

	b, err := json.MarshalIndent(accesstoken, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func cleanBase64(input string) string {
	re := regexp.MustCompile(`[^A-Za-z0-9+/=]`)
	cleaned := re.ReplaceAllString(input, "")

	padding := len(cleaned) % 4
	if padding > 0 {
		cleaned += strings.Repeat("=", 4-padding)
	}

	return cleaned
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	tokenCmd.Flags().StringP("parse", "p", "", "parse an agora token")
	tokenCmd.Flags().StringP("user", "u", "", "generate a new user token for use in SDK APIs")

}
