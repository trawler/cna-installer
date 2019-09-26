/*
Copyright © 2019 Karen Almog

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a remote backend for the Installer, if it does not exist.",
	Long: `
Create and generate the remote backend required for installation.
If you already have a remote backend, this step can be skipped.`,
	Run: func(cmd *cobra.Command, args []string) {
		initBackend()
	},
}

func initBackend() error {
	workDir := "data/terraform/tf-backend"
	tf, err := terraform.NewTerraformClient(workDir)
	if err != nil {
		return errors.New(fmt.Sprintln("Error occured", err))
	}

	initParams := terraform.NewTerraformInitParams()
	initParams.Opts()

	init := tf.Init(initParams)

	init.Initialise()
	init.Run()
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().Bool("remote-backend", true, "Enable terraform's Remote Backend")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
