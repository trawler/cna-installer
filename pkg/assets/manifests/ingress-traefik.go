package manifests

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
)

// traefikPolicyRules are the policies to give to the Traefik Ingress
var traefikClusterRoleName = "traefik-ingress-controller"
var traefikClusterBindingRoleName = "traefik-ingress-controller"
var namespace = "cna-installer"

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

func traefikClusterAuth(k8sClient *kubernetes.Clientset) {
	createClusterRole(k8sClient, traefikClusterRoleName, traefikPolicyRules)
	createClusterRoleBinding(
		k8sClient,
		traefikClusterBindingRoleName,
		traefikClusterRoleName,
		traefikClusterRoleName,
		namespace,
	)
}

// InstallTraefikIngressController is a general function that holds all the tasks for Installing Traefik
func InstallTraefikIngressController(k8sClient *kubernetes.Clientset) {
	traefikClusterAuth(k8sClient)
}
