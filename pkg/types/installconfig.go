package types

import (
	"fmt"

	"github.com/trawler/cna-installer/pkg/types/aws"
	"github.com/trawler/cna-installer/pkg/types/azure"
)

const (
	// InstallConfigVersion is the version supported by this package.
	// If you bump this, you must also update the list of convertable values in
	// pkg/types/conversion/installconfig.go
	InstallConfigVersion = "v1"
)

// InstallConfig is the configuration for an OpenShift install.
type InstallConfig struct {

	// SSHKey is the public ssh key to provide access to instances.
	// +optional
	SSHKey string `json:"sshKey,omitempty"`

	// BaseDomain is the base domain to which the cluster should belong.
	BaseDomain string `json:"baseDomain"`

	// Platform is the configuration for the specific platform upon which to
	// perform the installation.
	Platform `json:"platform"`
}

// ClusterDomain returns the DNS domain that all records for a cluster must belong to.
func (c *InstallConfig) ClusterDomain() string {
	return fmt.Sprintf("%s.%s", c.ObjectMeta.Name, c.BaseDomain)
}

// Platform is the configuration for the specific platform upon which to perform
// the installation. Only one of the platform configuration should be set.
type Platform struct {
	// AWS is the configuration used when installing on AWS.
	// +optional
	AWS *aws.Platform `json:"aws,omitempty"`

	// Azure is the configuration used when installing on Azure.
	// +optional
	Azure *azure.Platform `json:"azure,omitempty"`
}

// Name returns a string representation of the platform (e.g. "aws" if
// AWS is non-nil).  It returns an empty string if no platform is
// configured.
func (p *Platform) Name() string {
	switch {
	case p == nil:
		return ""
	case p.Azure != nil:
		return azure.Name
	default:
		return ""
	}
}
