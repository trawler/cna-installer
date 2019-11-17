package traefik

import (
	"github.com/trawler/cna-installer/pkg/assets/manifests/utils"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// ConfigData is a struct that holds the content of the ConfigMap yaml
var ConfigData = map[string]string{
	"data": ConfigMap,
}

// ConfigMap is the yaml contents of the configmap
var ConfigMap = `
data:
  traefik.toml: |
    # traefik.toml
    logLevel = "info"
    defaultEntryPoints = ["http","https"]
    [entryPoints]
      [entryPoints.http]
      address = ":80"
      compress = true
      [entryPoints.https]
      address = ":443"
      compress = true
        [entryPoints.https.tls]
          [[entryPoints.https.tls.certificates]]
          CertFile = "/ssl/tls.crt"
          KeyFile = "/ssl/tls.key"
      [entryPoints.traefik]
      address = ":8080"
    [ping]
    entryPoint = "http"
    [kubernetes]
    ingressClass = "traefik"
      [kubernetes.ingressEndpoint]
      publishedService = "ingress/ingress-traefik"
    [traefikLog]
      format = "json"
    [api]
      entryPoint = "traefik"
      dashboard = true`

var (
	AppName                = "traefik-ingress"
	Version                = "traefik:v2.0.4"
	ClusterBindingRoleName = "traefik-ingress-controller"
	ClusterRoleName        = "traefik-ingress-controller"
	ConfigMapName          = "traefik-ingress-controller"
	DeploymentName         = "traefik-ingress-controller"
	Namespace              = "cna-installer"
	ServiceAccountName     = "traefik-ingress-controller"

	// PolicyRules are the policies to give to the Traefik Ingress
	PolicyRules = []utils.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"services", "endpoints", "secrets"},
			Verbs:     []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{"extensions"},
			Resources: []string{"ingresses"},
			Verbs:     []string{"get", "list", "watch"},
		},
	}

	// ContainerSpec is an array of the containers to be deplooyed
	ContainerSpec = []utils.ContainerSpec{
		{
			Name:  AppName,
			Image: Version,
			Ports: []utils.ContainerPort{
				{Name: "http", ContainerPort: 80},
				{Name: "admin", ContainerPort: 8080},
			},
			Args: []string{"--configfile=/config/traefik.toml"},
		},
	}

	// WebUIServiceSpec service
	WebUIServiceSpec = utils.ServiceSpec{
		Ports: []utils.ServicePort{
			{
				Name:       "dashboard-http",
				Protocol:   "TCP",
				Port:       80,
				TargetPort: intstr.FromInt(8080),
			},
		},
		Selector: map[string]string{
			"app": AppName,
		},
		Protocol:        "TCP",
		SessionAffinity: "None",
		Type:            "ClusterIP",
	}
)
