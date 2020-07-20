package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Labels defines (JSON) structure of gitlab labels
type Labels []struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Color                  string `json:"color"`
	TextColor              string `json:"text_color"`
	Description            string `json:"description"`
	DescriptionHTML        string `json:"description_html"`
	OpenIssuesCount        int    `json:"open_issues_count"`
	ClosedIssuesCount      int    `json:"closed_issues_count"`
	OpenMergeRequestsCount int    `json:"open_merge_requests_count"`
	Subscribed             bool   `json:"subscribed"`
	Priority               int    `json:"priority"`
	IsProjectLabel         bool   `json:"is_project_label"`
}

var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "A brief description of your command",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(labelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
