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
	"os"
	"sort"

	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"
	"github.com/spf13/cobra"

	"text/tabwriter"
)

// buildListCmd represents the delete command
var buildListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists summarized details of all multiplayer server builds for a title.",
	Long:  `lists summarized details of all multiplayer server builds for a title.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := listBuilds()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	buildCmd.AddCommand(buildListCmd)
}

func listBuilds() error {
	settings := getSettings()
	entityToken := getEntityToken()
	listBuildSummariesData := &multiplayer.ListBuildSummariesRequestModel{}
	res5, err := multiplayer.ListBuildSummaries(settings, listBuildSummariesData, entityToken)
	if err != nil {
		return err
	}

	sort.Slice(res5.BuildSummaries, func(i, j int) bool {
		return res5.BuildSummaries[i].CreationTime.After(res5.BuildSummaries[j].CreationTime)
	})

	const padding = 3
	q := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintf(q, "BuildID\tBuildName\tCreationTime\t\n")
	for _, bs := range res5.BuildSummaries {
		fmt.Fprintf(q, "%s\t%s\t%s\n", bs.BuildId, bs.BuildName, bs.CreationTime)
	}
	q.Flush()

	return nil

}
