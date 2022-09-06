/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/EusRique/authentication/application/rest/routes"
	"github.com/spf13/cobra"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start rest server",
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		routes.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
	restCmd.Flags().StringP("port", "p", "3000", "REST Server port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
