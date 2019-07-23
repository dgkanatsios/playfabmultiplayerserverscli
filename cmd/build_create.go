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
		Name:     "MultiplayerServerCountPerVm",
		Validate: survey.Required,
		Prompt: &survey.Input{
			Message: "Number of servers per VM",
			Help:    "Number of servers per VM",
			Default: "1",
		},
	},
	{
		Name:     "BuildName",
		Validate: survey.Required,
		Prompt: &survey.Input{
			Message: "Build name",
			Help:    "Build name",
			Default: "MyCustomBuild",
		},
	},
	{
		Name:     "StartMultiplayerServerCommand",
		Validate: survey.Required,
		Prompt: &survey.Input{
			Message: "Start MultiplayerServer Command",
			Help:    "Start MultiplayerServer Command",
			Default: "C:\\Assets\\WindowsRunnerCSharp.exe",
		},
	},
}

var qsEx = []*survey.Question{
	{
		Name:     "VmSize",
		Validate: survey.Required,
		Prompt: &survey.Select{
			Message: "Virtual Machine Size",
			Help:    "Virtual Machine Size",
			Options: []string{
				string(multiplayer.AzureVmSizeStandard_D1_v2),
				string(multiplayer.AzureVmSizeStandard_D2_v2),
				string(multiplayer.AzureVmSizeStandard_D3_v2),
				string(multiplayer.AzureVmSizeStandard_D4_v2),
				string(multiplayer.AzureVmSizeStandard_D5_v2),
			},
		},
	},
}

var portQ = []*survey.Question{
	{
		Name:     "Name",
		Prompt:   &survey.Input{Message: "Port name?"},
		Validate: survey.Required,
	},
	{
		Name:     "Num",
		Prompt:   &survey.Input{Message: "Port number?"},
		Validate: survey.Required,
	},
	{
		Name: "Protocol",
		Prompt: &survey.Select{
			Message: "Port protocol?",
			Options: []string{
				string(multiplayer.ProtocolTypeTCP),
				string(multiplayer.ProtocolTypeUDP),
			}},
		Validate: survey.Required,
	},
}

var gameAssetQ = []*survey.Question{
	{
		Name:     "FileName",
		Prompt:   &survey.Input{Message: "Game asset Filename?", Default: "winrunnerSample.zip"},
		Validate: survey.Required,
	},
	{
		Name:     "MountPath",
		Prompt:   &survey.Input{Message: "Game asset Mountpath?", Default: "C:\\Assets\\"},
		Validate: survey.Required,
	},
}

var buildRegionQ = []*survey.Question{
	{
		Name:     "MaxServers",
		Prompt:   &survey.Input{Message: "MaxServers?", Default: "1"},
		Validate: survey.Required,
	},
	{
		Name:     "StandbyServers",
		Prompt:   &survey.Input{Message: "StandbyServers?", Default: "1"},
		Validate: survey.Required,
	},
	{
		Name: "Region",
		Prompt: &survey.Select{
			Message: "Region?",
			Options: []string{
				string(multiplayer.AzureRegionAustraliaEast),
				string(multiplayer.AzureRegionAustraliaSoutheast),
				string(multiplayer.AzureRegionBrazilSouth),
				string(multiplayer.AzureRegionCentralUs),
				string(multiplayer.AzureRegionEastAsia),
				string(multiplayer.AzureRegionEastUs),
				string(multiplayer.AzureRegionEastUs2),
				string(multiplayer.AzureRegionWestUs),
			}},
		Validate: survey.Required,
	},
}

func createBuild() error {
	data := &multiplayer.CreateBuildWithManagedContainerRequestModel{}

	err := survey.Ask(qs, data)
	if err != nil {
		return err
	}

	qsExRes := struct {
		VmSize string
	}{}

	err = survey.Ask(qsEx, &qsExRes)
	if err != nil {
		return err
	}

	data.VmSize = multiplayer.AzureVmSize(qsExRes.VmSize)

	data.Ports = []multiplayer.PortModel{}
	for {
		portEx := struct {
			Name     string
			Num      int32
			Protocol string
		}{}

		survey.Ask(portQ, &portEx)

		port := multiplayer.PortModel{}
		port.Name = portEx.Name
		port.Num = portEx.Num
		port.Protocol = multiplayer.ProtocolType(portEx.Protocol)
		data.Ports = append(data.Ports, port)

		addAnotherPort := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "Add another port?",
		}, &addAnotherPort); err != nil {
			return err
		}

		if !addAnotherPort {
			break
		}
	}

	data.GameAssetReferences = []multiplayer.AssetReferenceParamsModel{}
	for {
		gameAssetEx := multiplayer.AssetReferenceParamsModel{}
		survey.Ask(gameAssetQ, &gameAssetEx)

		data.GameAssetReferences = append(data.GameAssetReferences, gameAssetEx)

		addAnotherGameAssetReference := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "Add another game asset reference?",
		}, &addAnotherGameAssetReference); err != nil {
			return err
		}

		if !addAnotherGameAssetReference {
			break
		}
	}

	data.RegionConfigurations = []multiplayer.BuildRegionParamsModel{}
	for {
		buildRegionEx := struct {
			MaxServers     int32
			StandbyServers int32
			Region         string
		}{}

		survey.Ask(buildRegionQ, &buildRegionEx)

		buildRegion := multiplayer.BuildRegionParamsModel{}
		buildRegion.MaxServers = buildRegionEx.MaxServers
		buildRegion.StandbyServers = buildRegionEx.StandbyServers
		buildRegion.Region = multiplayer.AzureRegion(buildRegionEx.Region)
		data.RegionConfigurations = append(data.RegionConfigurations, buildRegion)

		addAnotherRegion := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "Add another region?",
		}, &addAnotherRegion); err != nil {
			return err
		}

		if !addAnotherRegion {
			break
		}
	}

	settings := getSettings()
	entityToken := getEntityToken()

	res4, err := multiplayer.CreateBuildWithManagedContainer(settings, data, entityToken)
	if err != nil {
		return err
	}
	fmt.Printf("%#v", res4)
	return nil
}
