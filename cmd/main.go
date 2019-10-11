/*
Copyright © 2019 Karen Almog

*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

var cfgFile string
var cluster *terraform.Cluster

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cna-installer",
	Short: "Creates a CNA cluster",
	Long: `cna-installer is a binary that installs, sets-up and configures a
kubernetes cluster with the CNA stack applications.`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cna-installer.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error

	// If no config file was provided, try to use the default config file.
	if cfgFile == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		cfgFile = path.Join(home, ".cna-installer.yaml")
	}

	// Parse the config file into a Cluster struct
	cluster, err = terraform.ParseConfigFile(cfgFile)
	if err != nil {
		os.Exit(1)
	}
}
