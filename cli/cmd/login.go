/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"box/helpers"
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to box app!",
	Long: `Login to box app example:

	box login --e email -p password`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		if email == "" {
			email = callEmail()
		}
		if password == "" {
			password = callPassword()
		}
		helpers.Login(email, password)
	},
}

// statusCmd represents the status of login
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "login status of box app!",
	Long: `login status of box app example:

	box login status`,
	Run: func(cmd *cobra.Command, args []string) {
		ls, email := helpers.LoginStatus()
		if ls {
			fmt.Println("Logged into Box App with " + email)
		} else {
			fmt.Println("Logged out of Box App")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().StringP("email", "e", "", "login email for Box App")
	loginCmd.Flags().StringP("password", "p", "", "login password for Box App")
	loginCmd.AddCommand(statusCmd)
}
