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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var titleID *string
var secretKey *string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if *titleID == "" || *secretKey == "" {
			fmt.Println("titleID and secretKey are required")
			return
		}

		token, err := login(*titleID, *secretKey)

		if err != nil {
			fmt.Println("Error logging in")
			return
		}

		viper.Set(titleIDConfig, titleID)
		viper.Set(entityTokenConfig, token)

		err = viper.WriteConfig()
		if err != nil {
			fmt.Printf("Cannot write config because of %s\n", err)
			return
		}
		fmt.Printf("Successfully logged in, config saved at %s\n", viper.ConfigFileUsed())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	titleID = loginCmd.Flags().StringP("title", "t", "", "Your Title ID")
	loginCmd.MarkFlagRequired("title")
	secretKey = loginCmd.Flags().StringP("secret", "s", "", "Your Title Secret")
	loginCmd.MarkFlagRequired("secret")
}
