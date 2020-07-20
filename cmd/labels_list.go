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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all labels from project or group.",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("GET", viper.GetString("baseURL")+"projects/"+viper.GetString("projectID")+"/labels?with_counts=true&per_page=-1", nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Private-Token", viper.GetString("PersonalAccessToken"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var labels Labels

		reqGet("projects/"+viper.GetString("projectID")+"/labels?with_counts=true", &labels)

		l, err := json.MarshalIndent(labels, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(l))

	},
}

func init() {
	labelsCmd.AddCommand(listCmd)
}
