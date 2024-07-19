/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/accesstoken2"
	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/chatTokenBuilder"
	apps "github.com/CarlsonYuan/agora-chat-cli/apps"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate and parse agora tokens",
	RunE: func(cmd *cobra.Command, args []string) error {

		expire := uint32(24 * 60 * 60)

		active, err := apps.GetApps(cmd).GetActiveAppOrExplicit(cmd)
		if err != nil {
			return err
		}

		fapp, _ := cmd.Flags().GetBool("app")
		if fapp {
			result, err := chatTokenBuilder.BuildChatAppToken(active.AppID, active.AppCertificate, expire)
			if err != nil {
				return err
			} else {
				cmd.Printf("token for chat app [%s]:\n%s\n", active.AppID, result)
			}
		}

		userID, _ := cmd.Flags().GetString("user")
		if userID != "" {
			userToken, err := chatTokenBuilder.BuildChatUserToken(active.AppID, active.AppCertificate, userID, expire)
			if err != nil {
				return err
			}
			cmd.Printf("Token for user [%s]:\n%s\n", userID, userToken)
		}

		token, _ := cmd.Flags().GetString("parse")
		if token != "" {
			version := token[:3]
			if version != "007" {
				cmd.PrintErrln("Not support, just for parsing token version 007!")
			}
			token := token[3:]
			accesstoken := accesstoken2.CreateAccessToken()
			finalToken := version + cleanBase64(token)
			result, err := accesstoken.Parse(finalToken)
			if !result {
				return err
			}
			b, err := json.MarshalIndent(accesstoken, "", "  ")
			if err != nil {
				return err
			}
			cmd.Printf("Token is %s\n\nToken information:\n%s\n", finalToken, string(b))
		}

		return nil
	},
}

func cleanBase64(input string) string {
	// Remove all non-base64 characters
	re := regexp.MustCompile(`[^A-Za-z0-9+/=]`)
	cleaned := re.ReplaceAllString(input, "")

	// Add padding if necessary
	padding := len(cleaned) % 4
	if padding > 0 {
		cleaned += strings.Repeat("=", 4-padding)
	}

	return cleaned
}
func init() {
	rootCmd.AddCommand(tokenCmd)

	tokenCmd.Flags().BoolP("app", "a", false, "generate a new chat application token for use in RESTful APIs")
	tokenCmd.Flags().StringP("parse", "p", "", "parse an agora token")
	tokenCmd.Flags().StringP("user", "u", "", "generate a new user token for use in SDK APIs")

}
