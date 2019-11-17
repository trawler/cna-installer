package manifests

import (
	"fmt"

	"github.com/trawler/cna-installer/pkg/assets/manifests/traefik"
	"github.com/trawler/cna-installer/pkg/assets/manifests/utils"
	"k8s.io/client-go/kubernetes"
)

// InstallTraefikIngressController is a general function that holds all the tasks for Installing Traefik
func InstallTraefikIngressController(k8sClient *kubernetes.Clientset) error {
	// Create ClusterRole
	if err := utils.CreateClusterRole(k8sClient, traefik.AppName, traefik.ClusterRoleName, traefik.PolicyRules); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ClusterRoleBinding
	if err := utils.CreateClusterRoleBinding(k8sClient, traefik.AppName, traefik.ClusterBindingRoleName, traefik.ClusterRoleName, traefik.ClusterRoleName, traefik.Namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ServiceAccount
	if err := utils.CreateServiceAccount(k8sClient, traefik.AppName, traefik.ServiceAccountName, traefik.Namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ConfigMap
	config, err := utils.ParseConfigData(traefik.ConfigMap)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	if err := utils.CreateConfigMap(k8sClient, traefik.AppName, traefik.ConfigMapName, config.Data, traefik.Namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create Deployment
	if err := utils.CreateDeployment(k8sClient, traefik.ConfigMapName, traefik.DeploymentName, traefik.ServiceAccountName, traefik.ContainerSpec, traefik.Namespace); err != nil {
		return fmt.Errorf("%v", err)
	}

	// web-ui service
	if err := utils.CreateService(k8sClient, traefik.AppName, "traefik-web-ui", traefik.WebUIServiceSpec, traefik.Namespace); err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
