/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
)

type sendOptions struct {
	// chat type
	IsPrivateChat bool
	IsGroupChat   bool
	IsRoomChat    bool

	Sender    string
	Receivers []string

	// message type
	IsTextMessage     bool
	IsCommandMessage  bool
	IsLocationMessage bool
	IsCustomMessage   bool

	// txt type
	Content string

	// cmd type
	Action string

	// loc type
	Latitude  string
	Longitude string
	Address   string

	// custom type
	CustomEvent string
	CustomExts  string
}

func messageCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "message",
		Short: "Manage messages",
	}
	cmd.AddCommand(sendCmd())
	return cmd
}

func sendCmd() *cobra.Command {
	opts := &sendOptions{}

	var cmd = &cobra.Command{
		Use:   "send",
		Short: "Send a message via Agora Chat Server",
		Example: heredoc.Doc(`
				$ agchat message send --users --text --content "Hello, world" --receivers demo_user_1,demo_user_2
				$ agchat message send --users --cmd --action "custom_action" --receivers demo_user_1,demo_user_2
				$ agchat message send --users --custom --custom-event "custom_event" --custom-exts '{"key1": "value1", "key2": "value2"}' --receivers demo_user_1,demo_user_2
				$ agchat message send --users --loc --lat "39.916668" --lon "116.383331" --addr "beijing" --receivers demo_user_1,demo_user_2
		`),
		RunE: func(cmd *cobra.Command, args []string) error {

			if !opts.IsPrivateChat && !opts.IsGroupChat && !opts.IsRoomChat {
				return fmt.Errorf("`--users`, `--groups` or `--rooms` required to send message")
			}
			if !opts.IsCommandMessage && !opts.IsTextMessage && !opts.IsCustomMessage && !opts.IsLocationMessage {
				return fmt.Errorf("`--text`, `--cmd`, `--custom` or `--loc` required to send message")
			}
			// message type
			var msg_type ac.MessageType
			var msg_body interface{}

			if opts.IsTextMessage {
				if opts.Content == "" {
					return fmt.Errorf("`--content` option is required when using `--text`")
				}
				msg_type = ac.MessageTypeText
				msg_body = ac.TextMessageBody{
					Msg: opts.Content,
				}
			} else if opts.IsCommandMessage {
				if opts.Action == "" {
					return fmt.Errorf("`--action` option is required when using `--cmd`")
				}
				msg_type = ac.MessageTypeCmd
				msg_body = ac.CMDMessageBody{
					Action: opts.Action,
				}
			} else if opts.IsCustomMessage {
				if opts.CustomEvent == "" {
					return fmt.Errorf("`--custom-event` option is required when using `--custom`")
				}
				msg_type = ac.MessageTypeCustom
				exts, err := parseJSONFlag(opts.CustomExts)
				if err != nil {
					return err
				}
				msg_body = ac.CustomMessageBody{
					CustomEvent: opts.CustomEvent,
					CustomExts:  exts,
				}
			} else if opts.IsLocationMessage {
				if opts.Latitude == "" && opts.Longitude == "" && opts.Address == "" {
					return fmt.Errorf("`--lat`, `--lon` and `--addr` options are required when using `--loc`")
				}
				msg_type = ac.MessageTypeLoc
				msg_body = ac.LocationMessageBody{
					Lat:  opts.Latitude,
					Lng:  opts.Longitude,
					Addr: opts.Address,
				}
			}
			message := ac.Message{
				From: opts.Sender,
				To:   opts.Receivers,
				Type: msg_type,
				Body: msg_body,
			}
			if opts.IsPrivateChat {
				msgIDs, err := client.Message().SendUsersMessage(&message)
				if err != nil {
					return err
				}
				for userID, msgID := range msgIDs {
					logger.Info("Successfully sent message", map[string]interface{}{
						"messageID":  msgID,
						"receiverID": userID,
					})
				}
			} else if opts.IsGroupChat {
				msgIDs, err := client.Message().SendGroupsMessage(&message)
				if err != nil {
					return err
				}
				for groupID, msgID := range msgIDs {
					logger.Info("Successfully sent message", map[string]interface{}{
						"messageID": msgID,
						"groupID":   groupID,
					})
				}
			} else if opts.IsRoomChat {
				msgIDs, err := client.Message().SendRoomsMessage(&message)
				if err != nil {
					return err
				}
				for roomID, msgID := range msgIDs {
					logger.Info("Successfully sent message", map[string]interface{}{
						"messageID": msgID,
						"roomID":    roomID,
					})
				}
			}

			return nil
		},
	}

	fl := cmd.Flags()

	// chat type
	fl.BoolVarP(&opts.IsPrivateChat, "users", "u", false, "Send a message to users")
	fl.BoolVarP(&opts.IsGroupChat, "groups", "g", false, "Send a message to groups")
	fl.BoolVarP(&opts.IsRoomChat, "rooms", "R", false, "Send a message to rooms")

	fl.StringVarP(&opts.Sender, "sender", "s", "", "Sender of the message (default 'admin')")
	fl.StringSliceVarP(&opts.Receivers, "receivers", "r", []string{}, "Receivers of the message")

	// message type
	fl.BoolVarP(&opts.IsTextMessage, "text", "", false, "Text type message")
	fl.BoolVarP(&opts.IsCommandMessage, "cmd", "", false, "Command type message")
	fl.BoolVarP(&opts.IsLocationMessage, "loc", "", false, "Location type message")
	fl.BoolVarP(&opts.IsCustomMessage, "custom", "", false, "Custom type message")

	// txt
	fl.StringVarP(&opts.Content, "content", "c", "", "The content of the text message")

	// cmd
	fl.StringVarP(&opts.Action, "action", "a", "", "The action of the command message")

	// loc
	fl.StringVarP(&opts.Latitude, "lat", "l", "", "The latitude of the location message")
	fl.StringVarP(&opts.Longitude, "lon", "L", "", "The longitude of the location message")
	fl.StringVarP(&opts.Address, "addr", "A", "", "The address of the location message")

	// custom
	fl.StringVarP(&opts.CustomEvent, "custom-event", "e", "", "The event of the custom message")
	fl.StringVarP(&opts.CustomExts, "custom-exts", "E", "", "The exts of the custom messagen JSON format (e.g., '{\"key1\": \"value1\", \"key2\": \"value2\"}')")

	return cmd
}

// Function to parse JSON string into map[string]string
func parseJSONFlag(jsonString string) (map[string]string, error) {
	result := make(map[string]string)

	if jsonString == "" {
		return result, nil
	}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
