/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/ycj3/agora-chat-cli/util"
)

type detailOptions struct {
	groupIDs []string
}

func groupCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "group",
		Short: "Manage groups",
	}
	cmd.AddCommand(detailCmd())
	return cmd
}

func detailCmd() *cobra.Command {
	opts := &detailOptions{}

	var cmd = &cobra.Command{
		Use:   "detail",
		Short: "Get the group detail info of the group(s)",
		Example: heredoc.Doc(`
				$ agchat group detail --groups <group-id-1>,<group-id-2>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			retust, err := client.Group().GetGroupDetail(opts.groupIDs)
			if err != nil {
				return err
			}
			util.Print(retust.Data, util.OutputFormatJSON, nil)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringSliceVarP(&opts.groupIDs, "groups", "g", []string{}, "ID of the groups")

	return cmd
}
