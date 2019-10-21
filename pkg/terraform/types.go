package terraform

import "os/exec"

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

// AzureAuth holds the required fields for Azure authentication
type AzureAuth struct {
	SubsctiptionID   string
	ClientID         string
	ClientSecret     string
	TenantID         string
	BackendAccessKey string
}

// TfInitParams is a struct that holds terraform init parameters
type TfInitParams struct {
	Backend       *bool
	BackendConfig string
	ForceCopy     bool
	FromModule    string
	Get           *bool
	GetPlugins    *bool
	Input         *bool
	Lock          *bool
	LockTimeout   int
	NoColor       bool
	PluginDir     string
	Reconfigure   bool
	Upgrade       *bool
	VerifyPlugins *bool
}

// TfPlanParams is a struct that holds terraform plan parameters
type TfPlanParams struct {
	AutoApprove      bool
	Destroy          bool
	DetailedExitcode bool
	Input            *bool
	Lock             *bool
	LockTimeout      int
	NoColor          bool
	ModuleDepth      *int
	Out              *string
	Parallelism      *int
	Refresh          *bool
	State            *string
	Target           []*string
	Var              map[string]string
	VarFile          []*string
}

// TFApply is a struct that holds terraform apply parameters
type TFApply struct {
	AutoApprove      *bool
	Destroy          bool
	DetailedExitcode bool
	Input            *bool
	Lock             *bool
	LockTimeout      int
	NoColor          bool
	ModuleDepth      *int
	Out              *string
	Parallelism      *int
	Refresh          *bool
	State            *string
	Target           []*string
	Var              map[string]string
	VarFile          []*string
}

// TfActionParams comment
type TfActionParams interface {
	Opts() map[string][]string
	OptsString() string
	OptsStringSlice() []string
}

// TfAction comment
type TfAction struct {
	action        string
	bin           *Executor
	Cmd           *exec.Cmd
	executionPath string
	opts          map[string]string
	params        TfActionParams
}
