package utils

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
