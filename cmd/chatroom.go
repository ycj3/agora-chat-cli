/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/ycj3/agora-chat-cli/util"
)

type roomDetailOptions struct {
	roomIDs []string
}

func roomCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "room",
		Short: "Manage chat rooms",
	}
	cmd.AddCommand(roomDetailCmd())
	return cmd
}

func roomDetailCmd() *cobra.Command {
	opts := &roomDetailOptions{}

	var cmd = &cobra.Command{
		Use:   "detail",
		Short: "Get the room detail info of the chat room(s)",
		Example: heredoc.Doc(`
				$ agchat room detail --rooms <chatroom-id-1>,<chatroom-id-2>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			retust, err := client.Room().GetChatroomDetail(opts.roomIDs)
			if err != nil {
				return err
			}
			util.Print(retust.Data, util.OutputFormatJSON, nil)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringSliceVarP(&opts.roomIDs, "rooms", "g", []string{}, "ID of the rooms")

	return cmd
}
