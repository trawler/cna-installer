package utils

import "k8s.io/apimachinery/pkg/util/intstr"

// ConfigMapData points to the yaml contents of the Data field of the configMap
type ConfigMapData struct {
	Data map[string]string `yaml:"data"`
}

// ContainerPort is a generic struct that hold the container ports info
type ContainerPort struct {
	Name          string
	ContainerPort int
}

// ContainerSpec is a generic struct that hold the container spec info
type ContainerSpec struct {
	Name  string
	Image string
	Ports []ContainerPort
	Args  []string
}

// PolicyRule is a generic struct that holds RBAC policy rules
type PolicyRule struct {
	Verbs           []string
	APIGroups       []string
	Resources       []string
	ResourceNames   []string
	NonResourceURLs []string
}

// ServicePort struct
type ServicePort struct {
	Name       string
	Port       int
	Protocol   string
	TargetPort intstr.IntOrString
}

// ServiceSpec struct
type ServiceSpec struct {
	Ports           []ServicePort
	Protocol        string
	Selector        map[string]string
	SessionAffinity string

	// ClusterIP / NodePort / LoadBalancer or ExternalName
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	Type string
}
