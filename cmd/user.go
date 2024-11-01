/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/ycj3/agora-chat-cli/util"
)

type userDetailOptions struct {
	uID string
}

type userCreateOptions struct {
	uID string
	pwd string
}

type userDeleteOptions struct {
	uID string
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
}

var onlineStatusCmd = &cobra.Command{
	Use:   "online-status",
	Short: "Query the online status of users",
	Long:  `Query the online status of users by specifying the users`,
	RunE: func(cmd *cobra.Command, args []string) error {

		users, _ := cmd.Flags().GetString("users")

		if users[len(users)-1] == ',' {
			logger.Warn("Extra spaces detected in --users flag. They will be removed.", nil)
			users = users[:len(users)-1]
		}

		userIDs := strings.Split(users, ",")

		statuses, err := client.User().UserOnlineStatuses(userIDs)

		if err != nil {
			return err
		}
		util.OutputJson(statuses)
		return nil
	},
}

func userDetailCmd() *cobra.Command {
	opts := &userDetailOptions{}

	var cmd = &cobra.Command{
		Use:   "detail",
		Short: "Get the detailed information of a specific user",
		Example: heredoc.Doc(`
				$ agchat user detail --user <user-id>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := client.User().QueryUser(opts.uID)
			if err != nil {
				return err
			}
			util.OutputJson(user)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.uID, "user", "u", "", "ID of the user")

	return cmd
}

func userCreateCmd() *cobra.Command {
	opts := &userCreateOptions{}

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		Example: heredoc.Doc(`
				$ agchat user create --user <user-id> --password <password>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := client.User().CreateUser(opts.uID, opts.pwd)
			if err != nil {
				return err
			}
			logger.Info("Successfully created user", map[string]interface{}{
				"uuid": user.UUID,
			})
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.uID, "user", "u", "", "ID of the user")
	fl.StringVarP(&opts.pwd, "password", "p", "", "Password of the user")

	return cmd
}

func userDeleteCmd() *cobra.Command {
	opts := &userDeleteOptions{}

	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a user",
		Example: heredoc.Doc(`
				$ agchat user delete --user <user-id>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := client.User().DeleteUser(opts.uID)
			if err != nil {
				return err
			}
			logger.Info(fmt.Sprintf("User: %s deleted", user.ID), nil)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.uID, "user", "u", "", "ID of the user")

	return cmd
}

func init() {

	userCmd.AddCommand(onlineStatusCmd)
	userCmd.AddCommand(userDetailCmd())
	userCmd.AddCommand(userCreateCmd())
	userCmd.AddCommand(userDeleteCmd())

	onlineStatusCmd.Flags().StringP("users", "u", "", "Comma-separated list of users to query the online status for")
	onlineStatusCmd.MarkFlagRequired("users")

}
