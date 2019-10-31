package terraform

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// DefaultCluster populates a default structure of a cluster
var defaultCluster = &Cluster{
	TfConfigVars: TfConfigVars{
		BaseDomain:     "",
		ClusterName:    "",
		ClusterOwner:   "",
		ClusterVersion: "1.14.3",
	},
	TfAzureVars: TfAzureVars{
		AgentCount:       3,
		AvailabilityZone: "francecentral",
		ClientID:         "",
		ClientSecret:     "",
		OSDisksize:       "30",
		OSType:           "Linux",
		VMSize:           "Standard_DS2_v2",
	},
}

// GetEnvVars get a cluster structure, and exports it
// to terraform environment variables
func GetEnvVars(c *Cluster) error {
	auth := new(AzureAuth)
	err := verifyAuth(auth)
	if err != nil {
		return fmt.Errorf("Failed to verify auth: %v", err)
	}

	if c.TfConfigVars.BaseDomain == "" {
		return fmt.Errorf("baseDomain is not set")
	}

	if c.TfConfigVars.ClusterName == "" {
		return fmt.Errorf("clusterName is not set")
	}

	if c.TfConfigVars.ClusterOwner == "" {
		return fmt.Errorf("clusterOwner is not set")
	}

	prefix := "TF_VAR"
	os.Setenv(fmt.Sprintf("%s_base_domain", prefix), c.TfConfigVars.BaseDomain)
	os.Setenv(fmt.Sprintf("%s_cluster_owner", prefix), c.TfConfigVars.ClusterOwner)
	os.Setenv(fmt.Sprintf("%s_k8s_cluster_name", prefix), c.TfConfigVars.ClusterName)
	os.Setenv(fmt.Sprintf("%s_k8s_version", prefix), c.TfConfigVars.ClusterVersion)

	os.Setenv(fmt.Sprintf("%s_agent_count", prefix), strconv.Itoa(c.TfAzureVars.AgentCount))
	os.Setenv(fmt.Sprintf("%s_agent_os_disk_size_gb", prefix), c.TfAzureVars.OSDisksize)
	os.Setenv(fmt.Sprintf("%s_agent_pool_name", prefix), c.TfAzureVars.AgentPoolName)
	os.Setenv(fmt.Sprintf("%s_agent_vm_size", prefix), c.TfAzureVars.VMSize)
	os.Setenv(fmt.Sprintf("%s_az_location", prefix), c.TfAzureVars.AvailabilityZone)
	os.Setenv(fmt.Sprintf("%s_azure_client_id", prefix), auth.ClientID)
	os.Setenv(fmt.Sprintf("%s_cluster_autoscaling", prefix),
		fmt.Sprintf("%t", c.TfAzureVars.AutoScaler))
	os.Setenv(fmt.Sprintf("%s_azure_client_secret", prefix), auth.ClientSecret)
	os.Setenv(fmt.Sprintf("%s_k8s_resource_group_name", prefix), c.TfAzureVars.ResourceGroup)
	os.Setenv(fmt.Sprintf("%s_public_key_file", prefix), c.TfAzureVars.PublicKeyFile)

	return nil
}

// ParseConfig parses the yaml file and merges it with the defaultCluster
// to return a Cluster struct
func ParseConfig(data []byte) (*Cluster, error) {
	cluster := defaultCluster

	// Read config file and merge default cluster with overrides
	if err := yaml.Unmarshal(data, &cluster); err != nil {
		return nil, fmt.Errorf("unable to parse yaml data file")
	}

	return cluster, nil
}

// ParseConfigFile parses a yaml file and returns, if successful, a Cluster.
func ParseConfigFile(path string) (*Cluster, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", path)
	}
	return ParseConfig(data)
}

func verifyAuth(a *AzureAuth) error {
	a.ClientID = os.Getenv("ARM_CLIENT_ID")
	if a.ClientID == "" {
		return fmt.Errorf("ARM_CLIENT_ID is not set")
	}

	a.ClientSecret = os.Getenv("ARM_CLIENT_SECRET")
	if a.ClientSecret == "" {
		return fmt.Errorf("ARM_CLIENT_SECRET is not set")
	}

	a.TenantID = os.Getenv("ARM_TENANT_ID")
	if a.TenantID == "" {
		return fmt.Errorf("ARM_TENANT_ID is not set")
	}

	a.SubsctiptionID = os.Getenv("ARM_SUBSCRIPTION_ID")
	if a.SubsctiptionID == "" {
		return fmt.Errorf("ARM_SUBSCRIPTION_ID is not set")
	}
	return nil
}
