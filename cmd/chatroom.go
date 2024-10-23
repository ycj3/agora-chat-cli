/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	agora_chat "github.com/ycj3/agora-chat-cli/agora-chat"
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
			result, err := client.Room().GetChatroomDetail(opts.roomIDs)
			if err != nil {
				if restErr, ok := err.(agora_chat.Error); ok {
					if restErr.ErrorType == "service_resource_not_found" {
						roomID := strings.Split(restErr.ErrorDescription, ":")[1]
						logger.Error("Failed to get chatroom details", map[string]interface{}{
							"error": fmt.Sprintf("The chatroom[ID: %s] is not found", roomID),
						})
					} else {
						logger.Error("Failed to get chatroom details", map[string]interface{}{
							"error": err.Error(),
						})
					}
				}
				return nil
			}
			util.Print(result.Data, util.OutputFormatJSON, nil)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringSliceVarP(&opts.roomIDs, "rooms", "g", []string{}, "ID of the rooms")

	return cmd
}
