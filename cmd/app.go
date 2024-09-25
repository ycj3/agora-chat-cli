package cmd

import (
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
	auth "github.com/ycj3/agora-chat-cli/agora-chat/auth"
	"github.com/ycj3/agora-chat-cli/cmdutil"
)

type createOption struct {
	WithEnvToken bool
	AppName      string
	BaseUrl      string
}

type removeOption struct {
	AppNames []string
}

type useOption struct {
	AppName string
}

func appCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Manage chat applications",
	}

	cmd.AddCommand(createCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(removeCmd())
	cmd.AddCommand(useCmd())

	cmdutil.DisableAuthCheck(cmd)

	return cmd
}

func useCmd() *cobra.Command {
	opts := &useOption{}
	var cmd = &cobra.Command{
		Use:   "use",
		Short: "Mark an application as the active one",
		Example: heredoc.Doc(`
				$ agchat app use --name <application-name>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			apps, err := acCfg.GetApps()
			if err != nil {
				return fmt.Errorf("failed to get apps: %s", err)
			}
			if err := apps.Use(opts.AppName); err != nil {
				return fmt.Errorf("failed to mark application: %s", err)
			}
			logger.Info("Successfully marked application as the active one", map[string]interface{}{
				"name": opts.AppName,
			})
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.AppName, "name", "n", "", "Name for the application")

	return cmd
}

func removeCmd() *cobra.Command {
	opts := &removeOption{}
	var cmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove one or more application",
		Example: heredoc.Doc(`
				$ agchat app remove --names <application-name1>,<application-name2>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			apps, err := acCfg.GetApps()
			if err != nil {
				return fmt.Errorf("failed to get apps: %s", err)
			}
			for _, appName := range opts.AppNames {
				if err := apps.Remove(appName); err != nil {
					return err
				}
				logger.Info("Successfully removed application", map[string]interface{}{
					"name": appName,
				})
			}
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringSliceVarP(&opts.AppNames, "names", "n", []string{}, "App names of the application")

	return cmd
}

func listCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list all chat applications",
		Example: heredoc.Doc(`
				$ agchat app list
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			apps, err := acCfg.GetApps()
			if err != nil {
				return fmt.Errorf("failed to get apps: %s", err)
			}
			return apps.ListAllApps()
		},
	}
	return cmd
}

func createCmd() *cobra.Command {
	opts := &createOption{}
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create an chat application using details from the Agora console",
		Example: heredoc.Doc(`
				$ agchat app create --env-token --name <application-name> --url <application-base-url>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			apps, err := acCfg.GetApps()
			if err != nil {
				return fmt.Errorf("failed to get apps: %s", err)
			}
			if opts.WithEnvToken {
				if !auth.HasEnvToken() {
					return errors.New("set the AC_TOKEN environment variable is required to create an App with the env token")
				}
				if opts.AppName == "" || opts.BaseUrl == "" {
					return fmt.Errorf("`--name` and `--url` are required to create an App with the env token")
				}
				if err := apps.Add(ac.App{Name: opts.AppName, BaseURL: opts.BaseUrl}); err != nil {
					return fmt.Errorf("failed to add application: %s", err)
				}
				logger.Info("Added application", map[string]interface{}{
					"name": opts.AppName,
				})
				return nil
			}
			return apps.RunQuestionnaire()
		},
	}

	fl := cmd.Flags()
	fl.StringVarP(&opts.AppName, "name", "n", "", "Name for the application")
	fl.BoolVarP(&opts.WithEnvToken, "env-token", "e", false, "Create an application with the env token")
	fl.StringVarP(&opts.BaseUrl, "url", "u", "", "Base URL of the application")
	return cmd
}
