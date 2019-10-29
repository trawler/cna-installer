package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/trawler/cna-installer/pkg/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var remoteBackend *terraform.AzureBackend

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
		clusterCreate()
	},
}

var clusterDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy a Kubernetes cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		clusterDestroy()
	},
}

func clusterCreate() {
	if err := initClusterWorkspace(); err != nil {
		fmt.Printf("\nERROR:\nfailed to initialize environment:\n\n ")
		log.Fatal(err)
	}
	if err := tfRun(); err != nil {
		fmt.Printf("\nERROR:\nfailed to run terraform:\n\n ")
		log.Fatal(err)
	}
}

func clusterDestroy() {
	if err := initClusterWorkspace(); err != nil {
		//fmt.Printf("failed to initialize environment: %v\n", err)
		log.Fatal(err)
	}
	if err := tfDestroy(); err != nil {
		fmt.Printf("Failed to clean up cluster:\n%v\n", err)
		log.Fatal(err)
	}
}

func initClusterWorkspace() error {
	logDir, err = getLogDir()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if err := getRemoteBackend(); err != nil {
		return fmt.Errorf("failed to configure remote backend: %v", err)
	}
	if err := prepareTFRun(); err != nil {
		return fmt.Errorf("failed to set terraform run environment: %v", err)
	}
	return nil
}

func prepareTFRun() error {
	stateFileName, err = getStateFilePath("aks")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Set the resource group name to `cluster name`-`cluster owner`-aks
	cluster.TfAzureVars.ResourceGroup = fmt.Sprintf(
		"%s-%s-aks",
		cluster.TfConfigVars.ClusterName,
		cluster.TfConfigVars.ClusterOwner,
	)

	// set cluster domain to `cluster name`.`cluster owner`.`base domain`
	cluster.TfConfigVars.BaseDomain = fmt.Sprintf("%s.%s.%s",
		cluster.TfConfigVars.ClusterName,
		cluster.TfConfigVars.ClusterOwner,
		cluster.TfConfigVars.BaseDomain,
	)

	// set the azure agent's pool name
	agentPoolName, err := sanitizeAgentPoolName(cluster.TfConfigVars.ClusterName, cluster.TfConfigVars.ClusterOwner)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	cluster.TfAzureVars.AgentPoolName = agentPoolName

	// Set the executionPath for the terraform backend config
	executionPath := "../../data/terraform/aks"
	tf, err = terraform.NewTerraformClient(executionPath, logDir)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	initParams.Backend = terraform.TruePtr()
	initParams.BackendConfig = append([]string{
		fmt.Sprintf("storage_account_name=%s", remoteBackend.StorageAccountName),
		fmt.Sprintf("container_name=%s", remoteBackend.ContainerName),
		fmt.Sprintf("key=%s", remoteBackend.Key),
		fmt.Sprintf("access_key=%s", remoteBackend.AccessKey),
	})
	// set the location of the state File
	planParams.State = &stateFileName
	return nil
}

func getRemoteBackend() error {
	stateFileName, _ = getStateFilePath("backend")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	state, err := terraform.ReadStateFile(stateFileName)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	accessKey, err := getRemoteBackendAccessKey(state)
	if err != nil {
		return fmt.Errorf("cannot fetch remote access key: %v", err)
	}
	containerName, err := getStorageContainerName(state)
	if err != nil {
		return fmt.Errorf("cannot fetch storage container name: %v", err)
	}

	// build remote backend auth struct
	remoteBackend = &terraform.AzureBackend{
		AccessKey:          accessKey,
		ContainerName:      "cna-tfstate",
		Key:                fmt.Sprintf("terraform.%s.tfstate", cluster.TfConfigVars.ClusterOwner),
		StorageAccountName: containerName,
	}
	return nil
}

func getRemoteBackendAccessKey(tfstate *terraform.State) (string, error) {
	backend, err := terraform.LookupResource(tfstate, "", "azurerm_storage_account", "tf-cna-backend")
	if err != nil {
		return "", fmt.Errorf("failed to lookup remote backend: %v", err)
	}
	if len(backend.Instances) == 0 {
		return "", fmt.Errorf("no remote backend instance found")
	}
	accessKey, _, err := unstructured.NestedString(backend.Instances[0].Attributes, "primary_access_key")
	if err != nil {
		return "", fmt.Errorf("no primary access key found for remote backend")
	}
	return accessKey, nil
}

func getStorageContainerName(tfstate *terraform.State) (string, error) {
	storageContainer, err := terraform.LookupResource(tfstate, "", "azurerm_storage_container", "tf-storage-container")
	if err != nil {
		return "", fmt.Errorf("failed to lookup storage container: %v", err)
	}
	if len(storageContainer.Instances) == 0 {
		return "", fmt.Errorf("no storage container instance found")
	}
	containerName, _, err := unstructured.NestedString(storageContainer.Instances[0].Attributes, "storage_account_name")
	if err != nil {
		return "", fmt.Errorf("no container name found for storage container")
	}
	return containerName, nil
}

func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.AddCommand(clusterCreateCmd)
	clusterCmd.AddCommand(clusterDestroyCmd)
}
