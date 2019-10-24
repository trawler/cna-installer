package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
)

// createCmd represents the create command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a Cluster",
	Long:  `Create or Destroy an AKS or AWS cluster.`,
}

var clusterCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Kubernetes cluster.",
	Long: `
Creats a Kubernetes cluster. In the case of Azure, AKS will be used.
If you already have a remote backend, make sure the access details are stated in cna-installer.yaml.
Otherwise, a new remote backend will be created and used.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initClusterWorkspace(); err != nil {
			fmt.Printf("failed to initialize environment: %v\n", err)
			os.Exit(1)
		}
	},
}

func initClusterWorkspace() error {
	// Populate TF_VAR Environment variables
	if err = terraform.GetEnvVars(cluster); err != nil {
		return fmt.Errorf("%v", err)
	}

	// debugging:
	fmt.Printf("DEBUG:\n cluster: %+v\n\n", cluster)

	logDir, err = getLogDir()
	stateFileName, _ = getStateFilePath("backend")

	state, err := terraform.ReadStateFile(stateFileName)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// debugging:
	fmt.Printf("DEBUG::\n state: %v", state)

	return nil
}

func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.AddCommand(clusterCreateCmd)
}
