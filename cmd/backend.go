package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

var initParams = terraform.NewTerraformInitParams()
var planParams = terraform.NewTerraformPlanParams()

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
		if err := initWorkspace(); err != nil {
			fmt.Printf("failed to initialize environment: %v\n", err)
			os.Exit(1)
		}
		if err := initBackend(); err != nil {
			fmt.Printf("Error initializing backend:\n%v\n", err)
			os.Exit(1)
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
		if err := initWorkspace(); err != nil {
			fmt.Printf("failed to initialize environment: %v\n", err)
			os.Exit(1)
		}
		if err := destroyBackend(); err != nil {
			fmt.Printf("Error destroying backend:\n%v\n", err)
			os.Exit(1)
		}
	},
}

func initWorkspace() error {
	// Get the logDir path and direct tf output to state file
	logDir, err = getLogDir()
	stateFileName, _ = getStateFilePath("backend")

	// Populate TF_VAR Environment variables
	if err = terraform.GetEnvVars(cluster); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Set the executionPath for the terraform backend config
	executionPath := "../../data/terraform/tf-backend"
	tf, err = terraform.NewTerraformClient(executionPath, logDir)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// set the location of the state File
	planParams.State = &stateFileName

	return nil
}

func initBackend() error {
	// Get Init opts
	initParams.Opts()

	// Run terraform init
	init := tf.Init(initParams)

	init.Initialise()
	init.Run()

	// Run terraform plan
	planParams.Opts()
	plan := tf.Plan(planParams)
	plan.Initialise()

	if err = plan.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Run terraform apply
	apply := tf.Apply(planParams)
	apply.Initialise()

	if err = apply.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func destroyBackend() error {
	destroy := tf.Destroy(planParams)
	destroy.Initialise()

	if err = destroy.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(backendCmd)
	backendCmd.AddCommand(backendInitCmd)
	backendCmd.AddCommand(backendDestroyCmd)
}
