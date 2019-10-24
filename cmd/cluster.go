package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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

	logDir, err = getLogDir()
	stateFileName, _ = getStateFilePath("backend")

	state, err := terraform.ReadStateFile(stateFileName)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	accessKey, err := getRemoteBackendAccessKey(state)
	if err != nil {
		return fmt.Errorf("cannot fetch remote access key: %v", err)
	}
	aKeyEncrypted, _ := base64.URLEncoding.DecodeString(accessKey)
	fmt.Printf("Found access key in state file: %v\n", string(aKeyEncrypted))

	return nil
}

func getRemoteBackendAccessKey(tfstate *terraform.State) (string, error) {
	tfBackend, err := terraform.LookupResource(tfstate, "", "azurerm_storage_account", "tf-backend")
	if err != nil {
		return "", errors.Wrap(err, "failed to lookup remote backend")
	}
	if len(tfBackend.Instances) == 0 {
		return "", errors.New("no remote backend instance found")
	}
	accessKey, _, err := unstructured.NestedString(tfBackend.Instances[0].Attributes, "primary_access_key")
	if err != nil {
		return "", errors.New("no primary access key found for remote backend")
	}
	return accessKey, nil
}
func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.AddCommand(clusterCreateCmd)
}
