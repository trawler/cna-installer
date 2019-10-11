package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

var err error

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a remote backend for the Installer, if it does not exist.",
	Long: `
Create and generate the remote backend required for installation.
If you already have a remote backend, this step can be skipped.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := initBackend()
		if err != nil {
			fmt.Printf("Error initializing backend:\n%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().Bool("remote-backend", true, "Create terraform's Remote Backend")
}

func initBackend() error {
	if cluster.TfAzureVars.ClientID != "" && cluster.TfAzureVars.ClientSecret != "" {
		return fmt.Errorf("clientID AND clientSecret seem to be configured.\nAre you sure you need to set-up a new one? ")
	}

	// Populate TF_VAR Environment variables
	err = terraform.GetEnvVars(cluster)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	workDir := "data/terraform/tf-backend"
	tf, err := terraform.NewTerraformClient(workDir)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	initParams := terraform.NewTerraformInitParams()
	initParams.Opts()

	// Run terraform init
	init := tf.Init(initParams)
	init.Initialise()
	init.Run()

	// Run terraform plan

	planParams := terraform.NewTerraformPlanParams()
	planParams.Opts()
	plan := tf.Plan(planParams)
	plan.Initialise()
	plan.Run()
	return nil
}
