package manifests

import (
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	namespace                     = "cna-installer"
	serviceAccountName            = "traefik-ingress-controller"
	traefikClusterBindingRoleName = "traefik-ingress-controller"
	traefikClusterRoleName        = "traefik-ingress-controller"
)

// traefikPolicyRules are the policies to give to the Traefik Ingress
var traefikPolicyRules = []rbacv1.PolicyRule{
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

// InstallTraefikIngressController is a general function that holds all the tasks for Installing Traefik
func InstallTraefikIngressController(k8sClient *kubernetes.Clientset) error {
	// Create ClusterRole
	if err := createClusterRole(k8sClient, traefikClusterRoleName, traefikPolicyRules); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ClusterRoleBinding
	err := createClusterRoleBinding(
		k8sClient,
		traefikClusterBindingRoleName,
		traefikClusterRoleName,
		traefikClusterRoleName,
		namespace,
	)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Create ServiceAccount
	if err := createServiceAccount(k8sClient, serviceAccountName, namespace); err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
