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

// buildCreateCmd represents the create command
var buildCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createBuild()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	buildCmd.AddCommand(buildCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var buildName *string
var ports *string
var assetReferences *string
var startCommand *string

func createBuild() error {
	settings, entityToken := getLoginSettings()
	createBuildData := &multiplayer.CreateBuildWithManagedContainerRequestModel{
		MultiplayerServerCountPerVm: 1,
	}
	createBuildData.Ports = []multiplayer.PortModel{multiplayer.PortModel{
		Name:     "game_port",
		Num:      3600,
		Protocol: multiplayer.ProtocolTypeTCP,
	}}
	createBuildData.BuildName = "golangTest"
	createBuildData.GameAssetReferences = []multiplayer.AssetReferenceParamsModel{
		multiplayer.AssetReferenceParamsModel{
			FileName:  "winrunnerSample.zip",
			MountPath: "C:\\Assets\\",
		},
	}
	createBuildData.RegionConfigurations = []multiplayer.BuildRegionParamsModel{
		multiplayer.BuildRegionParamsModel{
			MaxServers:     1,
			StandbyServers: 1,
			Region:         multiplayer.AzureRegionEastUs,
		},
	}
	createBuildData.StartMultiplayerServerCommand = "C:\\Assets\\WindowsRunnerCSharp.exe"
	createBuildData.VmSize = multiplayer.AzureVmSizeStandard_D2_v2
	res4, err := multiplayer.CreateBuildWithManagedContainer(settings, createBuildData, entityToken)
	if err != nil {
		return err
	}
	log.Printf("%#v", res4)
	return nil
}
