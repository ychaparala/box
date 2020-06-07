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
	"bufio"
	"fmt"
	"html"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "singup for box app!",
	Long: `You can signup for box app For example:

box signup --e email -p password`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		if email == "" || !helpers.ValidateEmail(email) {
			for true {
				email = callEmail()
				if helpers.ValidateEmail(email) {
					break
				}
			}
		}
		if password == "" || !helpers.ValidatePassword(password) {
			for true {
				password = callPassword()
				if helpers.ValidatePassword(password) {
					if password == callConfirmPassword() {
						break
					} else {
						fmt.Println("Passwords didnt match")
					}
				}
			}
		}
		// Register User
		if helpers.SignUP(email, password) {
			fmt.Println("Welcome to Box App " + email + " " + html.UnescapeString("&#"+strconv.Itoa(128075)+";"))
		}
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	signupCmd.Flags().StringP("email", "e", "", "signup email for Box App")
	signupCmd.Flags().StringP("password", "p", "", "signup password for Box App")
}

func callEmail() (email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")
	return
}

func callPassword() string {
	fmt.Print("Enter password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	return string(bytePassword)
}

func callConfirmPassword() string {
	fmt.Print("Confirm password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	return string(bytePassword)
}
