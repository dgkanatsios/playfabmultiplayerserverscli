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
	"log"

	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"
	"github.com/spf13/cobra"
)

// serverEnableCmd represents the request command
var serverEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables Multiplayer servers functionality for a title",
	Long:  `Enables Multiplayer servers functionality for a title`,
	Run: func(cmd *cobra.Command, args []string) {
		err := enableMultiplayerServersForTitle()
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	serverCmd.AddCommand(serverEnableCmd)
}

func enableMultiplayerServersForTitle() error {
	settings := getSettings()
	entityToken := getEntityToken()
	enableMultiplayerServersForTitleRequest := &multiplayer.EnableMultiplayerServersForTitleRequestModel{}
	res2, err := multiplayer.EnableMultiplayerServersForTitle(settings, enableMultiplayerServersForTitleRequest, entityToken)
	if err != nil {
		return err
	}

	log.Printf("Multiplayer status for title %s is %s", settings.TitleId, res2.Status)
	return nil
}
