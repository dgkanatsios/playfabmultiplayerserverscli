package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"
)

// deleteCmd represents the delete command
var deleteAssetCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a multiplayer server game asset for a title",
	Long:  `deletes a multiplayer server game asset for a title.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteAsset(*assetToDelete)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var assetToDelete *string

func init() {
	assetCmd.AddCommand(deleteAssetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	assetToDelete = deleteAssetCmd.Flags().StringP("asset", "f", "", "Asset filename to delete")
}

func deleteAsset(asset string) error {
	if asset == "" {
		log.Fatal("Asset filename cannot be empty")
	}
	settings, entityToken := getLoginSettings()
	deleteAssetRequest := &multiplayer.DeleteAssetRequestModel{FileName: asset}
	_, err := multiplayer.DeleteAsset(settings, deleteAssetRequest, entityToken)
	if err != nil {
		return err
	}
	return nil
}
