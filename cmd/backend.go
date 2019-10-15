package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

var stateFileName string

// backendCmd represents the backend command
var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Manage the remote backend",
	Long: `
Create or destroy a remote backend, where the terraform
state is saved. If you already have a remote backend set up, this step
can be skipped.`,
}

var backendInitCmd = &cobra.Command{
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

var backendDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backend init called")
	},
}

func init() {
	rootCmd.AddCommand(backendCmd)
	backendCmd.AddCommand(backendInitCmd)
	backendCmd.AddCommand(backendDestroyCmd)
}

func initBackend() error {
	if cluster.TfAzureVars.ClientID != "" && cluster.TfAzureVars.ClientSecret != "" {
		return fmt.Errorf("clientID AND clientSecret seem to be configured.\nAre you sure you need to set-up a new one? ")
	}
	// Get the logDir path and direct tf output to state file
	logDir, err = getLogDir()
	stateFileName, _ = stateFile("backend")

	// Populate TF_VAR Environment variables
	err = terraform.GetEnvVars(cluster)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Set the executionPath for the terraform backend config
	executionPath := "../data/terraform/tf-backend"
	tf, err := terraform.NewTerraformClient(executionPath, logDir)
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
	planParams.State = &stateFileName

	plan := tf.Plan(planParams)
	plan.Initialise()
	err = plan.Run()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Run terraform apply
	apply := tf.Apply(planParams)
	apply.Initialise()
	apply.Run()

	return nil
}
