package azure

import (
	"encoding/json"
)

// Auth is the collection of credentials that will be used by terrform.
type Auth struct {
	ClientID     string `json:"azure_client_id,omitempty"`
	ClientSecret string `json:"azure_client_secret,omitempty"`
}

type Agent struct {
	Count      int    `json:"count"`
	DiskSizeGB int    `json:"osDiskSizeGB"`
	OSType     string `json:"osType,omitempty"`
	VmSize     string `json:"vmSize,omitempty"`
}

type config struct {
	Auth                        `json:",inline"`
	Agent                       `json:",inline"`
	AvailabilityZone            string `json:"az_location,omitempty"`
	BaseDomainResourceGroupName string `json:"azure_base_domain_resource_group_name,omitempty"`
}

// TFVars generates Azure-specific Terraform variables launching the cluster.
func TFVars(
	auth Auth,
	agent Agent,
	az string,
) ([]byte, error) {
	cfg := &config{
		Auth:             auth,
		AvailabilityZone: az,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
