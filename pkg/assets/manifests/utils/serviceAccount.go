package utils

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateServiceAccount comment
func CreateServiceAccount(k8sClient *kubernetes.Clientset,
	appName string,
	serviceAccountName string,
	customNamespace string) error {

	labels := map[string]string{
		"app": appName,
	}

	serviceAccount := corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceAccountName,
			Namespace: customNamespace,
			Labels:    labels,
		},
	}

	client := k8sClient.CoreV1().ServiceAccounts(customNamespace)
	_, err := client.Create(&serviceAccount)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create ServiceAccount %s: %v", serviceAccountName, err)
		}
		_, err = client.Update(&serviceAccount)
		if err != nil {
			return fmt.Errorf("Failed to update ServiceAccount %q: %v", serviceAccountName, err)
		}
		fmt.Printf("ServiceAccount %q updated\n", serviceAccountName)
	} else {
		fmt.Printf("ServiceAccount %q created\n", serviceAccountName)
	}
	return nil
}
