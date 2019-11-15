package utils

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateNamespace creates the 'cna-installer' namespace
func CreateNamespace(k8sClient *kubernetes.Clientset, customNamespace string) error {
	namespace := corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind: "Namespace",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: customNamespace,
		},
	}

	_, err := k8sClient.CoreV1().Namespaces().Create(&namespace)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create namespace %s: %v", customNamespace, err)
		}
		fmt.Printf("Namespace %q already exists\n", customNamespace)
		return nil
	}
	fmt.Printf("Namespace %q created\n", customNamespace)
	return nil
}
