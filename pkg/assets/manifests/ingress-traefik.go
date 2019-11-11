package manifests

import (
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
func InstallTraefikIngressController(k8sClient *kubernetes.Clientset) {
	createClusterRole(k8sClient, traefikClusterRoleName, traefikPolicyRules)
	createClusterRoleBinding(
		k8sClient,
		traefikClusterBindingRoleName,
		traefikClusterRoleName,
		traefikClusterRoleName,
		namespace,
	)
	createServiceAccount(k8sClient, serviceAccountName, namespace)
}
