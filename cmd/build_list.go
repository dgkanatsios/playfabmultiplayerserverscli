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

// buildListCmd represents the delete command
var buildListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := listBuilds()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	buildCmd.AddCommand(buildListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listBuilds() error {
	settings, entityToken := getLoginSettings()
	listBuildSummariesData := &multiplayer.ListBuildSummariesRequestModel{}
	res5, err := multiplayer.ListBuildSummaries(settings, listBuildSummariesData, entityToken)
	if err != nil {
		return err
	}

	fmt.Println("BuildID, BuildName, CreationTime")
	for _, bs := range res5.BuildSummaries {
		fmt.Printf("%s, %s, %s\n", bs.BuildId, bs.BuildName, bs.CreationTime)
	}
	return nil

}
