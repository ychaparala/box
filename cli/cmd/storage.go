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
	// "box/helpers"
	"fmt"

	"github.com/spf13/cobra"
)

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storage called")
	},
}

// statusCmd represents the storage command
var availableCmd = &cobra.Command{
	Use:   "available",
	Short: "available storage of box app storage",
	Long: `available storage of box app storage For example:

	box storage available`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("IN")
		// row := helpers.GetUsageToBoxApp()
		// const col = 50
		// Clear the screen by printing \x0c.
		// fmt.Println(row)
		// bar := fmt.Sprintf("\x0c[%%-%vs]", col)
		// fmt.Printf(bar, strings.Repeat("=", int(math.Ceil(float64(row["usage"])*50/float64(row["limit"]))))+html.UnescapeString("&#"+strconv.Itoa(9889)+";"))
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.AddCommand(availableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
