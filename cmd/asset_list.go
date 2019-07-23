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

// assetListCmd represents the delete command
var assetListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists summarized details of all multiplayer assets for a title.",
	Long:  `lists summarized details of all multiplayer assets for a title.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := listAssets()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	assetCmd.AddCommand(assetListCmd)
}

func listAssets() error {
	settings := getSettings()
	entityToken := getEntityToken()

	st := ""
	const padding = 3
	q := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintf(q, "FileName\t\n")

	responses := []multiplayer.AssetSummaryModel{}
	i := 0
	for {
		listAssetSummariesData := &multiplayer.ListAssetSummariesRequestModel{PageSize: 10, SkipToken: st}
		res5, err := multiplayer.ListAssetSummaries(settings, listAssetSummariesData, entityToken)
		if err != nil {
			return err
		}

		if res5.SkipToken == "" {
			break
		}

		st = res5.SkipToken
		responses = append(responses, res5.AssetSummaries...)
		i+=int(res5.PageSize)
		fmt.Printf("Downloaded %d asset summaries\n", i)
	}

	sort.Slice(responses, func(i, j int) bool {
		return responses[i].FileName < responses[j].FileName
	})
	for _, bs := range responses {
		fmt.Fprintf(q, "%s\t\n", bs.FileName)
	}
	q.Flush()

	return nil

}
