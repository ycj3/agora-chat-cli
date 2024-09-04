/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ac "github.com/ycj3/agora-chat-cli/agora-chat"
)

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage push providers added to an application",
}
var listProvidersCmd = &cobra.Command{
	Use:   "list",
	Short: "list all providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := ac.NewClient()
		res, err := client.Provider().ListPushProviders()
		if err != nil {
			return fmt.Errorf("failed to list providers: %w", err)
		}

		if len(res.Entities) == 0 {
			fmt.Println("no provider added")
			return nil
		}
		logger.Info("", map[string]interface{}{
			"count":    len(res.Entities),
			"provides": res.Entities,
		})
		return nil
	},
}

var deleteProviderCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid, _ := cmd.Flags().GetString("uuid")
		client := ac.NewClient()
		res, err := client.Provider().DeletePushProvider(uuid)
		if err != nil {
			return fmt.Errorf("failed to delete the provider: %w", err)
		}

		msg := fmt.Sprintf("Push provider [%s] was deleted successfully", res.Entities[len(res.Entities)-1].Name)
		logger.Info(msg, nil)
		return nil
	},
}

// APNS P8
var apnsCmd = &cobra.Command{
	Use:   "insert-apns",
	Short: "Insert an APNS push provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		apnsTeamId, _ := cmd.Flags().GetString("team-id")
		apnsKeyId, _ := cmd.Flags().GetString("key-id")

		apns := getPushProviderFromCommonFlags(cmd)

		apns.Type = ac.PushProviderAPNS
		apns.ApnsPushSettings = &ac.APNSConfig{
			TeamId: apnsTeamId,
			KeyId:  apnsKeyId,
		}
		client := ac.NewClient()
		res, err := client.Provider().InsertPushProvider(apns)
		if err != nil {
			logger.Error("failed to send request,", map[string]interface{}{
				"error": err,
			})
			return nil
		}
		logger.Info("Push provider inserted successfully", map[string]interface{}{
			"results": res.Entities,
		})
		return nil
	},
}

var fcmCmd = &cobra.Command{
	Use:   "insert-fcm",
	Short: "Insert an FCM push provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		fcmPushType, _ := cmd.Flags().GetString("push-type")
		fcmPriority, _ := cmd.Flags().GetString("priority")
		fcmProjectId, _ := cmd.Flags().GetString("project-id")
		fcmVersion, _ := cmd.Flags().GetString("version")

		fcm := getPushProviderFromCommonFlags(cmd)
		fcm.Type = ac.PushProviderFCM
		fcm.FcmPushSettings = &ac.FCMConfig{
			PushType:  fcmPushType,
			Priority:  fcmPriority,
			ProjectId: fcmProjectId,
			Version:   fcmVersion,
		}
		client := ac.NewClient()
		res, err := client.Provider().InsertPushProvider(fcm)
		if err != nil {
			logger.Error("failed to send request,", map[string]interface{}{
				"error": err,
			})
			return nil
		}
		logger.Info("Push provider inserted successfully", map[string]interface{}{
			"results": res.Entities,
		})
		return nil
	},
}

var huaweiCmd = &cobra.Command{
	Use:   "insert-huawei",
	Short: "Insert a Huawei push provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		category, _ := cmd.Flags().GetString("category")
		activityClass, _ := cmd.Flags().GetString("activityClass")

		hw := getPushProviderFromCommonFlags(cmd)
		hw.Type = ac.PushProviderFCM
		hw.HuaweiPushSettings = &ac.HuaweiConfig{
			Category:      category,
			ActivityClass: activityClass,
		}
		client := ac.NewClient()
		res, err := client.Provider().InsertPushProvider(hw)
		if err != nil {
			logger.Error("failed to send request,", map[string]interface{}{
				"error": err,
			})
			return nil
		}
		logger.Info("Push provider inserted successfully", map[string]interface{}{
			"results": res.Entities,
		})
		return nil
	},
}

func init() {
	pushCmd.AddCommand(providerCmd)
	providerCmd.AddCommand(listProvidersCmd)

	providerCmd.AddCommand(deleteProviderCmd)
	deleteProviderCmd.Flags().StringP("uuid", "u", "", "the uuid of the provide")

	// ANPS
	providerCmd.AddCommand(apnsCmd)
	addCommonFlags(apnsCmd)
	// Provider-specific flags for APNS
	apnsCmd.Flags().String("team-id", "", "team ID")
	apnsCmd.MarkFlagRequired("team-id")
	apnsCmd.Flags().String("key-id", "", "key ID")
	apnsCmd.MarkFlagRequired("key-id")

	//FCM
	providerCmd.AddCommand(fcmCmd)
	addCommonFlags(fcmCmd)
	// Provider-specific flags for FCM
	fcmCmd.Flags().String("push-type", "", "push type")
	fcmCmd.Flags().String("priority", "", "priority")
	fcmCmd.Flags().String("project-id", "", "project ID")
	fcmCmd.Flags().String("version", "", "version")

	providerCmd.AddCommand(huaweiCmd)
	addCommonFlags(huaweiCmd)
	// Provider-specific flags for HUAWEI
	huaweiCmd.Flags().String("category", "", "category")
	huaweiCmd.Flags().String("activityClass", "", "activityClass")

}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Name of the Certificate")
	cmd.Flags().String("env", "", "Environment (e.g., PRODUCTION, DEVELOPMENT)")
	cmd.Flags().String("cert", "", "Certificate path")
	cmd.Flags().String("pkg", "", "Package name")
}

func getPushProviderFromCommonFlags(cmd *cobra.Command) ac.PushProvider {
	name, _ := cmd.Flags().GetString("name")
	env, _ := cmd.Flags().GetString("env")

	cert, _ := cmd.Flags().GetString("cert")
	pkg, _ := cmd.Flags().GetString("pkg")

	// Create the provider object based on the flags
	return ac.PushProvider{
		Name:        name,
		Env:         env,
		Certificate: cert,
		PackageName: pkg,
	}
}
