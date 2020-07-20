package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var foo string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "A brief descriptReadInConfigion of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("token", "t", "", "Gitlab API token")
	rootCmd.PersistentFlags().IntP("project", "p", 0000, "Project ID")
	rootCmd.PersistentFlags().BoolP("group", "g", false, "Perform actions on top-level group.")

	rootCmd.MarkPersistentFlagRequired("project")

	viper.BindPFlag("projectID", rootCmd.PersistentFlags().Lookup("project"))
	viper.BindPFlag("PersonalAccessToken", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("group", rootCmd.PersistentFlags().Lookup("group"))

}

// initConfig reads in config
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".gitlab" (without extension).
	viper.SetConfigName(".gitlab")
	viper.SetConfigType("toml")
	viper.AddConfigPath(home)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file: ", viper.ConfigFileUsed())
	} else {
		fmt.Println("msg", "error reading config", "filename", viper.ConfigFileUsed(), "err", err)
		os.Exit(1)
	}
}
