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

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

// buildCreateCmd represents the create command
var buildCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a build",
	Long:  `creates a PlayFab Multiplayer Servers build`,
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

var qs = []*survey.Question{
	{
		Name: "serversPerVM",
		Prompt: &survey.Input{
			Message: "Number of servers per VM",
			Help:    "Number of servers per VM",
			Default: "1",
		},
	},
	{
		Name: "buildName",
		Prompt: &survey.Input{
			Message: "Build name",
			Help:    "Build name",
			Default: "MyCustomBuild",
		},
	},
	{
		Name: "startCommand",
		Prompt: &survey.Input{
			Message: "Start MultiplayerServer Command",
			Help:    "Start MultiplayerServer Command",
			Default: "C:\\Assets\\WindowsRunnerCSharp.exe",
		},
	},
	{
		Name: "startCommand",
		Prompt: &survey.Select{
			Message: "Virtual Machine Size",
			Help:    "Virtual Machine Size",
			Options: []string{"Standard_D1_v2", "Standard_D2_v2", "Standard_D3_v2", "Standard_D4_v2", "Standard_D5_v2"},
		},
	},
}

func createBuild() error {
	settings := getSettings()
	entityToken := getEntityToken()
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
			FileName:  "winrunnerSample6.zip",
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
