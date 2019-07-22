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

// buildGetCmd represents the delete command
var buildGetCmd = &cobra.Command{
	Use:   "get",
	Short: "lists summarized details of all multiplayer server builds for a title.",
	Long:  `lists summarized details of all multiplayer server builds for a title.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := getBuild()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var buildIDToGet *string

func init() {
	buildCmd.AddCommand(buildGetCmd)

	buildIDToGet = buildGetCmd.Flags().StringP("buildID", "b", "", "BuildID of the Build to be retrieved")
	buildGetCmd.MarkFlagRequired("buildID")
}

func getBuild() error {
	settings := getSettings()
	entityToken := getEntityToken()
	getBuildData := &multiplayer.GetBuildRequestModel{BuildId: *buildIDToGet}
	res5, err := multiplayer.GetBuild(settings, getBuildData, entityToken)
	if err != nil {
		return err
	}

	fmt.Println(prettyPrint(res5))

	return nil

}
