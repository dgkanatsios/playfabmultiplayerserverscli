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

// buildDeleteCmd represents the delete command
var buildDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteBuild()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	buildCmd.AddCommand(buildDeleteCmd)

	buildIDToDelete = buildDeleteCmd.Flags().StringP("buildID", "b", "", "BuildID of the Build to be deleted")
	buildDeleteCmd.MarkFlagRequired("buildID")
}

var buildIDToDelete *string

func deleteBuild() error {
	settings := getSettings()
	entityToken := getEntityToken()
	if *buildIDToDelete == "" {
		return fmt.Errorf("BuildID cannot be empty")
	}
	deleteBuildData := &multiplayer.DeleteBuildRequestModel{BuildId: *buildIDToDelete}
	_, err := multiplayer.DeleteBuild(settings, deleteBuildData, entityToken)
	if err != nil {
		return err
	}
	return nil
}
