package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete labels from project or group.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose which labels to delete:")

		fmt.Println(viper.AllSettings())

		var labels Labels
		reqGet("projects/"+viper.GetString("projectID")+"/labels?with_counts=true", &labels)

		var ids []string
		for i := range labels {
			ids = append(ids, labels[i].Name)
		}
		ids = append(ids, "DELETE ALL LABELS")

		prompt := promptui.Select{
			Label: "",
			Items: ids,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result != "DELETE ALL LABELS" {
			fmt.Println("Deleting label: " + result)
			req("DELETE", "projects/"+viper.GetString("projectID")+"/labels/"+result)
		} else {
			fmt.Println("Delete all labels?")
			if confirm() == true {
				fmt.Println("Your party! Deleting all labels")

				for i := range labels {
					req("DELETE", "projects/"+viper.GetString("projectID")+"/labels/"+strconv.Itoa((labels[i].ID)))
				}
			}
		}

	},
}

func init() {
	labelsCmd.AddCommand(deleteCmd)
}
