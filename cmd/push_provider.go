/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
	"github.com/ycj3/agora-chat-cli/util"
)

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage push providers added to an application",
}
var listProvidersCmd = &cobra.Command{
	Use:   "list",
	Short: "list all providers",
	RunE: func(cmd *cobra.Command, args []string) error {

		res, err := client.Provider().ListPushProviders()
		if err != nil {
			return fmt.Errorf("failed to list providers: %w", err)
		}

		if len(res.Entities) == 0 {
			fmt.Println("no provider added")
			return nil
		}
		util.Print(res.Entities, util.OutputFormatJSON, nil)
		return nil
	},
}

type deleteOptions struct {
	Uuids []string
}

func deleteProviderCmd() *cobra.Command {
	opts := &deleteOptions{}

	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete the provider",
		Example: heredoc.Doc(`
				$ agchat push provider delete --uuid <provider-uuid1>,<provider-uuid2>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, uuid := range opts.Uuids {
				res, err := client.Provider().DeletePushProvider(uuid)
				if err != nil {
					return fmt.Errorf("failed to delete the provider: %w", err)
				}

				logger.Info("Deleted push provider", map[string]interface{}{
					"name": res.Entities[0].Name,
				})
			}
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringSliceVarP(&opts.Uuids, "uuid", "u", nil, "Uuids of the push provider")

	return cmd
}

type upsertOption struct {

	// PushProvide type
	IsAPNs   bool
	IsFCM    bool
	IsHuawei bool

	NotifierId string // Same as UUID, it is used when editing the notifier.

	IsEnvDev bool

	Name        string
	Environment string
	Certificate string
	PackageName string
	Sound       string

	CertificateFilePath string
	CertificatePassword string
	// *** FCM ***
	// push message type
	IsBoth         bool
	IsNotifacation bool
	IsData         bool

	// Push Priority
	IsPriorityHigh   bool
	IsPriorityNormal bool

	SupportAPNs bool

	// *** APNs ***
	TeamID string
	KeyID  string
}

func upsertFcmCmd() *cobra.Command {
	opts := &upsertOption{}

	var cmd = &cobra.Command{
		Use:   "upsert",
		Short: "Inserts a push provider if it doesn't exist, or updates the existing push provider if it does",
		Example: heredoc.Doc(`
				# Insert a APNs (P12) push provider
				$ agchat push provider upsert --apns --name <certificate-name> --package-name <bundle-id> --dev-env --file <apns-p12-file> --password <password>

				# Insert a APNs (P8) push provider
				$ agchat push provider upsert --apns --name <certificate-name> --package-name <bundle-id> --dev-env --key-id <key-id> --team-id <team-id> --file <apns-p8-file>

				# Insert a FCM (V1) push provider
				$ agchat push provider upsert --fcm --name <fcm-sender-id> --file <fcm-service-account-json-file-path> --high --both

				# Enable APNs specific fields (e.g. sound, mutable-content)
				$ agchat push provider upsert --fcm --support-apns --nofitier-id <notifitier-id>

				# Disable APNs specific fields (e.g. sound, mutable-content)
				$ agchat push provider upsert --fcm --support-apns=false --nofitier-id <notifitier-id>
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			// var providerType ac.PushProviderType
			var env ac.EnvironmentType
			var provider ac.PushProvider

			if !opts.IsEnvDev {
				env = ac.EnvProduct
			} else {
				env = ac.EnvDevelopment
			}

			if opts.IsFCM {
				var pushType string
				var priority string

				if opts.IsNotifacation {
					pushType = ac.FCMPushNotification
				} else if opts.IsData {
					pushType = ac.FCMPushData
				} else {
					pushType = ac.FCMPushBoth
				}

				if opts.IsPriorityHigh {
					priority = ac.FCMPushPriorityHigh
				} else {
					priority = ac.FCMPushPriorityNormal
				}

				if opts.NotifierId != "" {
					provider = ac.PushProvider{
						NotifierId: opts.NotifierId,
						FcmPushSettings: &ac.FCMConfig{
							SupportAPNs: opts.SupportAPNs,
						},
					}
				} else {
					provider = ac.PushProvider{
						Name:     opts.Name,
						Provider: ac.PushProviderFCM,
						Env:      ac.EnvProduct,
						File:     opts.CertificateFilePath,
						FcmPushSettings: &ac.FCMConfig{
							PushType: pushType,
							Priority: priority,
							Version:  "v1",
							Sound:    opts.Sound,
						},
					}
				}
			} else if opts.IsAPNs {
				provider = ac.PushProvider{
					Name:        opts.Name,
					PackageName: opts.PackageName,
					Provider:    ac.PushProviderAPNS,
					Env:         env,
					File:        opts.CertificateFilePath,
					Passphrase:  opts.CertificatePassword,
					ApnsPushSettings: &ac.APNSConfig{
						TeamId: opts.TeamID,
						KeyId:  opts.KeyID,
						Sound:  opts.Sound,
					},
				}
			} else if opts.IsHuawei {
				return fmt.Errorf("not support yet for Huawei")
			} else {
				return fmt.Errorf("unsupported push provider")
			}
			res, err := client.Provider().UpsertPushProvider(provider)
			if err != nil {
				logger.Error("failed to upsert push provider request,", map[string]interface{}{
					"error": err,
				})
				return nil
			}
			logger.Info("Successfully upserted push provider", map[string]interface{}{
				"uuid": res.Entities[0].UUID,
			})
			return nil
		},
	}
	fl := cmd.Flags()

	// provider type
	fl.BoolVarP(&opts.IsAPNs, "apns", "", false, "APNs")
	fl.BoolVarP(&opts.IsFCM, "fcm", "", false, "FCM")
	fl.BoolVarP(&opts.IsHuawei, "huawei", "", false, "Huawei")

	// base config
	fl.StringVarP(&opts.Name, "name", "n", "", "Name of the Certificate")
	fl.BoolVarP(&opts.IsEnvDev, "dev-env", "e", false, "Set to true for DEVELOPMENT environment, false for PRODUCTION")
	fl.StringVarP(&opts.Certificate, "cert", "c", "", "Certificate path")
	fl.StringVarP(&opts.PackageName, "package-name", "p", "", "Package name")
	fl.StringVarP(&opts.Sound, "sound", "s", "default", "Sound")

	fl.StringVarP(&opts.CertificateFilePath, "file", "f", "", "Path of the certificate file")
	fl.StringVarP(&opts.CertificatePassword, "password", "P", "", "Passwrod of the certificate file")

	fl.StringVarP(&opts.NotifierId, "notifier-id", "N", "", "Same as UUID for the push provider")

	// *** FCM ***
	// push message type
	fl.BoolVarP(&opts.IsBoth, "both", "", false, "Both")
	fl.BoolVarP(&opts.IsNotifacation, "noti", "", false, "Notification type")
	fl.BoolVarP(&opts.IsData, "data", "", false, "Data type")

	// Push Priority
	fl.BoolVarP(&opts.IsPriorityHigh, "high", "", false, "Priority high")
	fl.BoolVarP(&opts.IsPriorityNormal, "normal", "", false, "Priority normal")

	fl.BoolVarP(&opts.SupportAPNs, "support-apns", "S", false, "Whether support APNs specific field (e.g. sound, mutable-content)")

	// *** APNs ***
	fl.StringVarP(&opts.TeamID, "team-id", "t", "", "Team ID of the APNs Certificate")
	fl.StringVarP(&opts.KeyID, "key-id", "k", "", "Key ID of the APNs Certificate")

	return cmd
}

// var huaweiCmd = &cobra.Command{
// 	Use:   "insert-huawei",
// 	Short: "Insert a Huawei push provider",
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		category, _ := cmd.Flags().GetString("category")
// 		activityClass, _ := cmd.Flags().GetString("activityClass")

// 		hw := getPushProviderFromCommonFlags(cmd)
// 		hw.Type = ac.PushProviderFCM
// 		hw.HuaweiPushSettings = &ac.HuaweiConfig{
// 			Category:      category,
// 			ActivityClass: activityClass,
// 		}

// 		res, err := client.Provider().InsertPushProvider(hw)
// 		if err != nil {
// 			logger.Error("failed to send request,", map[string]interface{}{
// 				"error": err,
// 			})
// 			return nil
// 		}
// 		logger.Info("Push provider inserted successfully", map[string]interface{}{
// 			"results": res.Entities,
// 		})
// 		return nil
// 	},
// }

func init() {
	pushCmd.AddCommand(providerCmd)

	providerCmd.AddCommand(listProvidersCmd)
	providerCmd.AddCommand(upsertFcmCmd())
	providerCmd.AddCommand(deleteProviderCmd())

	// providerCmd.AddCommand(huaweiCmd)
	// addCommonFlags(huaweiCmd)
	// Provider-specific flags for HUAWEI
	// huaweiCmd.Flags().String("category", "", "category")
	// huaweiCmd.Flags().String("activityClass", "", "activityClass")

}
