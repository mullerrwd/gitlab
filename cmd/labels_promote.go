package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// promoteCmd represents the promote command
var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote project labels to group labels.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose which labels to promote:")

		var labels Labels
		reqGet("projects/"+viper.GetString("projectID")+"/labels?per_page=-1", &labels)

		var ids []string
		for i := range labels {
			ids = append(ids, labels[i].Name)
		}
		ids = append(ids, "Promote all labels")

		prompt := promptui.Select{
			Label: "",
			Items: ids,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result != "Promote all labels" {
			fmt.Println("Promoting label: " + result)
			req("PUT", "projects/"+viper.GetString("projectID")+"/labels/"+result+"/promote")
		} else {
			fmt.Println("Are you sure to promote all labels to group labels?")
			if confirm() == true {
				for i := range labels {
					req("PUT", "projects/"+viper.GetString("projectID")+"/labels/"+strconv.Itoa((labels[i].ID))+"/promote")
				}
			}
		}
	},
}

func init() {
	labelsCmd.AddCommand(promoteCmd)
}
