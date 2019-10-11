package terraform

import (
	"fmt"
	"io/ioutil"
	"os"

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
	if c.TfConfigVars.BaseDomain == "" {
		return fmt.Errorf("baseDomain is not set")
	}

	if c.TfConfigVars.ClusterName == "" {
		return fmt.Errorf("clusterName is not set")
	}

	if c.TfConfigVars.ClusterOwner == "" {
		return fmt.Errorf("clusterOwner is not set")
	}

	c.TfAzureVars.ResourceGroup = fmt.Sprintf("%s_%s", c.TfConfigVars.ClusterName, c.TfConfigVars.ClusterOwner)

	prefix := "TF_VAR"
	os.Setenv(fmt.Sprintf("%s_base_domain", prefix), c.TfConfigVars.BaseDomain)
	os.Setenv(fmt.Sprintf("%s_dns_prefix", prefix), c.TfConfigVars.ClusterOwner)
	os.Setenv(fmt.Sprintf("%s_k8s_cluster_name", prefix), c.TfConfigVars.ClusterName)
	os.Setenv(fmt.Sprintf("%s_k8s_resource_group_name", prefix), c.TfAzureVars.ResourceGroup)

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
