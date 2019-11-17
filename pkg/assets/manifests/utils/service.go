package utils

import (
	"fmt"

	"k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateService creates a k8s service
func CreateService(k8sClient *kubernetes.Clientset,
	appName string,
	serviceName string,
	serviceSpec ServiceSpec,
	customNamespace string) error {

	labels := map[string]string{
		"app": appName,
	}

	spec, err := generateServiceSpec(serviceSpec)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	service := corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind: "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: customNamespace,
			Labels:    labels,
		},
		Spec: *spec,
	}

	client := k8sClient.CoreV1().Services(customNamespace)
	_, err = client.Create(&service)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create Service %s: %v", serviceName, err)
		}
		fmt.Printf("Service %q already exists\n", serviceName)
	} else {
		fmt.Printf("Service %q created\n", serviceName)
	}
	return nil
}
