package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add labels to project or group.",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {

		type Labelsadd []struct {
			Name        string `json:"name"`
			Color       string `json:"color"`
			Board       bool   `json:"board"`
			Priority    int    `json:"priority"`
			Description string `json:"description,omitempty"`
		}

		var labels Labelsadd

		jsonFile, err := os.Open(viper.GetString("config"))
		if err != nil {
			fmt.Println(err)
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &labels)
		for i := range labels {
			var label = "name=" + labels[i].Name +
				"&color=" + labels[i].Color +
				"&board=" + strconv.FormatBool(labels[i].Board) +
				"&priority=" + strconv.Itoa(labels[i].Priority) +
				"&description=" + labels[i].Description
			reqPOST("projects/"+viper.GetString("projectID")+"/labels/", label)
		}
	},
}

func init() {
	labelsCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("config", "c", "", "The config file containing the declaration of labels to be added. E.g: labels.json")
	addCmd.MarkFlagRequired("config")
	viper.BindPFlag("config", addCmd.Flags().Lookup("config"))
}
