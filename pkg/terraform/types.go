package terraform

// Cluster defines the config for a cluster.
type Cluster struct {
	TfAzureVars  TfAzureVars  `json:",inline" yaml:"azure,omitempty"`
	TfConfigVars TfConfigVars `json:",inline" yaml:"config,omitempty"`
}

// TfConfigVars comment
type TfConfigVars struct {
	BaseDomain     string `json:",inline" yaml:"baseDomain,omitempty"`
	ClusterName    string `json:",inline" yaml:"clusterName,omitempty"`
	ClusterOwner   string `json:",inline" yaml:"clusterOwner,omitempty"`
	ClusterVersion string `json:",inline" yaml:"clusterVersion,omitempty"`
}

// TfAzureVars comment
type TfAzureVars struct {
	AgentCount       int    `json:",inline" yaml:"agentCount,omitempty"`
	AvailabilityZone string `json:",inline" yaml:"availabilityZone,omitempty"`
	ClientID         string `json:",inline" yaml:"clientID,omitempty"`
	ClientSecret     string `json:",inline" yaml:"clientSecret,omitempty"`
	OSDisksize       string `json:",inline" yaml:"agentOSDiskSizeGB,omitempty"`
	OSType           string `json:",inline" yaml:"agentOSType,omitempty"`
	ResourceGroup    string
	VMSize           string `json:",inline" yaml:"agentVMSize,omitempty"`
}
