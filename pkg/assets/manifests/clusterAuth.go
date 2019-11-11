package manifests

import (
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateClusterRole creates a k8s Cluster Role
func createClusterRole(
	k8sClient *kubernetes.Clientset,
	clusterRoleName string,
	rules []rbacv1.PolicyRule,
) error {
	clusterRole := rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterRoleName,
		},
		Rules: rules,
	}

	client := k8sClient.RbacV1().ClusterRoles()
	_, err := client.Create(&clusterRole)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create ClusterRole %q: %v", clusterRoleName, err)
		}
		_, err = client.Update(&clusterRole)
		if err != nil {
			return fmt.Errorf("Failed to update ClusterRole %q: %v", clusterRoleName, err)
		}
		fmt.Printf("ClusterRole %q updated\n", clusterRoleName)
	} else {
		fmt.Printf("ClusterRole %q created\n", clusterRoleName)
	}
	return nil
}

// CreateClusterRoleBinding creates a k8s Cluster Role Binding
func createClusterRoleBinding(
	k8sClient *kubernetes.Clientset,
	clusterBindingRoleName string,
	serviceAccountName,
	clusterRoleName string,
	namespace string,
) error {
	roleBinding := rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterBindingRoleName,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     clusterRoleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      serviceAccountName,
				Namespace: namespace,
			},
		},
	}
	_, err := k8sClient.RbacV1().ClusterRoleBindings().Create(&roleBinding)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create ClusterRoleBinding %s: %v", clusterBindingRoleName, err)
		}
		fmt.Printf("ClusterRoleBinding %q already exists\n", clusterBindingRoleName)
		return nil
	}
	fmt.Printf("ClusterRoleBinding %q created, bound %q to %q\n", clusterBindingRoleName, serviceAccountName, clusterRoleName)
	return nil
}
