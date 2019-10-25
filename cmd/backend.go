package cmd

import (
	"fmt"
	"log"

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
		backendInit()
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
		backendDestroy()
	},
}

func backendInit() {
	if err := initWorkspace(); err != nil {
		fmt.Printf("failed to initialize environment: %v\n", err)
		log.Fatal(err)
	}
	if err := tfRun(); err != nil {
		fmt.Printf("Error initializing backend:\n%v\n", err)
		log.Fatal(err)
	}
}

func backendDestroy() {
	if err := initWorkspace(); err != nil {
		fmt.Printf("failed to initialize environment: %v\n", err)
		log.Fatal(err)
	}
	if err := tfDestroy(); err != nil {
		fmt.Printf("Failed to clean up cluster backend:\n%v\n", err)
		log.Fatal(err)
	}
}

func initWorkspace() error {
	// Get the logDir path and direct tf output to state file
	logDir, err = getLogDir()
	stateFileName, _ = getStateFilePath("backend")

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

func init() {
	rootCmd.AddCommand(backendCmd)
	backendCmd.AddCommand(backendInitCmd)
	backendCmd.AddCommand(backendDestroyCmd)
}
