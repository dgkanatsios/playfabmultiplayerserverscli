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
	Short: "creates an asset for a multiplayer server title",
	Long:  `creates an asset for a multiplayer server title`,
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

	assetToCreate = createAssetCmd.Flags().StringP("asset", "f", "", "Absolute filename of asset to create")
	assetCmd.MarkFlagRequired("asset")
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
	settings := getSettings()
	entityToken := getEntityToken()
	getAssetUploadURLRequest := &multiplayer.GetAssetUploadUrlRequestModel{FileName: fileName}

	res3, err := multiplayer.GetAssetUploadUrl(settings, getAssetUploadURLRequest, entityToken)
	if err != nil {
		return err
	}

	file, err := os.Open(asset)
	if err != nil {
		return err
	}
	credential := azblob.NewAnonymousCredential()
	assetURL, err := url.Parse(res3.AssetUploadUrl)
	if err != nil {
		return err
	}
	url := azblob.NewBlockBlobURL(*assetURL, azblob.NewPipeline(credential, azblob.PipelineOptions{}))
	_, err = azblob.UploadFileToBlockBlob(context.Background(), file, url, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16,
	})
	if err != nil {
		return err
	}
	return nil
}
