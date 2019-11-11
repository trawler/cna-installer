package manifests

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createServiceAccount(k8sClient *kubernetes.Clientset,
	serviceAccountName string,
	customNamespace string) error {

	serviceAccount := corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceAccountName,
			Namespace: customNamespace,
		},
	}

	//_, err := k8sClient.CoreV1().Namespaces().Create(&namespace)
	_, err := k8sClient.CoreV1().ServiceAccounts(customNamespace).Create(&serviceAccount)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create ServiceAccount %s: %v", customNamespace, err)
		}
		fmt.Printf("ServiceAccount %q already exists\n", serviceAccountName)
		return nil
	}
	fmt.Printf("ServiceAccount %q created\n", serviceAccountName)
	return nil
}
