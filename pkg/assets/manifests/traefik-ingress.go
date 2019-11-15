package manifests

import (
	"fmt"

	"github.com/trawler/cna-installer/pkg/assets/manifests/utils"
	"k8s.io/client-go/kubernetes"
)

var (
	traefikAppName = "traefik-ingress"
	traefikVersion = "traefik:v2.0.4"

	// traefikPolicyRules are the policies to give to the Traefik Ingress
	traefikPolicyRules = []utils.PolicyRule{
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

	// Array of deployed containers
	traefikContainerSpec = []utils.ContainerSpec{
		{
			Name:  traefikAppName,
			Image: traefikVersion,
			Ports: []utils.ContainerPort{
				{
					Name:          "http",
					ContainerPort: 80,
				},
				{
					Name:          "admin",
					ContainerPort: 8080,
				},
			},
			Args: []string{"--configfile=/config/traefik.toml"},
		},
	}
)

// InstallTraefikIngressController is a general function that holds all the tasks for Installing Traefik
func InstallTraefikIngressController(k8sClient *kubernetes.Clientset) error {
	var (
		clusterBindingRoleName = "traefik-ingress-controller"
		clusterRoleName        = "traefik-ingress-controller"
		configMapName          = "traefik-ingress-controller"
		deploymentName         = "traefik-ingress-controller"
		namespace              = "cna-installer"
		serviceAccountName     = "traefik-ingress-controller"
	)

	// Create ClusterRole
	if err := utils.CreateClusterRole(k8sClient, traefikAppName, clusterRoleName, traefikPolicyRules); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ClusterRoleBinding
	if err := utils.CreateClusterRoleBinding(k8sClient, traefikAppName, clusterBindingRoleName, clusterRoleName, clusterRoleName, namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ServiceAccount
	if err := utils.CreateServiceAccount(k8sClient, traefikAppName, serviceAccountName, namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ConfigMap
	config, err := utils.ParseConfigData(configData)
	if err != nil {
		// debug
	}

	if err := utils.CreateConfigMap(k8sClient, traefikAppName, configMapName, config.Data, namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create Deployment
	if err := utils.CreateDeployment(k8sClient, configMapName, deploymentName, serviceAccountName, traefikContainerSpec, namespace); err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
