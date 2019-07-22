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
	"fmt"
	"log"

	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"
	"github.com/spf13/cobra"
)

// serverRequestCmd represents the request command
var serverRequestCmd = &cobra.Command{
	Use:   "request",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := requestServer(*buildIDRequest)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	serverCmd.AddCommand(serverRequestCmd)

	buildIDRequest = serverRequestCmd.Flags().StringP("buildID", "b", "", "BuildID of the Build to be requested")
}

var buildIDRequest *string

func requestServer(buildID string) error {
	settings := getSettings()
	entityToken := getEntityToken()
	if buildID == "" {
		return fmt.Errorf("BuildID cannot be empty")
	}
	requestMultiplayerServerDetails := &multiplayer.RequestMultiplayerServerRequestModel{}
	requestMultiplayerServerDetails.PreferredRegions = []multiplayer.AzureRegion{multiplayer.AzureRegionEastUs}
	requestMultiplayerServerDetails.BuildId = buildID
	requestMultiplayerServerDetails.SessionId = "00000000-0000-0000-0000-000000000001"
	requestMultiplayerServerDetails.SessionCookie = "test cookie"
	res7, err := multiplayer.RequestMultiplayerServer(settings, requestMultiplayerServerDetails, entityToken)
	if err != nil {
		return err
	}
	log.Printf("%#v", res7)
	return nil
}
