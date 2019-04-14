// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"

	"github.com/spf13/cobra"
)

// createAssetCmd represents the create command
var createAssetCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createAsset(*assetToCreate)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var assetToCreate *string

func init() {
	assetCmd.AddCommand(createAssetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	assetToCreate = createAssetCmd.Flags().StringP("asset", "f", "", "Absolute filename of asset to create")
}

func createAsset(asset string) error {
	if asset == "" {
		log.Fatal("Asset path cannot be empty")
	}
	if !filepath.IsAbs(asset) {
		log.Fatal("Asset path should be absolute")
	}
	// get the filename
	_, fileName := filepath.Split(asset)
	settings, entityToken := getLoginSettings()
	getAssetUploadUrlRequest := &multiplayer.GetAssetUploadUrlRequestModel{FileName: fileName}
	//"C:\\projects\\playfabmultiplayerserverscli\\winrunnerSample.zip"
	//"/mnt/c/projects/playfabmultiplayerserverscli/winrunnerSample.zip"}
	res3, err := multiplayer.GetAssetUploadUrl(settings, getAssetUploadUrlRequest, entityToken)
	if err != nil {
		return err
	}
	log.Printf("%#v", res3)

	file, err := os.Open(res3.FileName)
	if err != nil {
		return err
	}
	credential := azblob.NewAnonymousCredential()
	assetUrl, err := url.Parse(res3.AssetUploadUrl)
	if err != nil {
		return err
	}
	url := azblob.NewBlockBlobURL(*assetUrl, azblob.NewPipeline(credential, azblob.PipelineOptions{}))
	_, err = azblob.UploadFileToBlockBlob(context.Background(), file, url, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16,
	})
	if err != nil {
		return err
	}
	return nil
}
