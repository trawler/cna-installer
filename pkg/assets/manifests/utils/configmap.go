package utils

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateConfigMap comment
func CreateConfigMap(k8sClient *kubernetes.Clientset,
	appName string,
	configMapName string,
	configMapData map[string]string,
	namespace string) error {

	labels := map[string]string{
		"app": appName,
	}

	configMap := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appName,
			Namespace: namespace,
			Labels:    labels,
		},
		Data: configMapData,
	}

	//	client := k8sClient.AppsV1().Deployments(namespace)
	client := k8sClient.CoreV1().ConfigMaps(namespace)
	_, err := client.Create(&configMap)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create ConfigMap %s: %v", configMapName, err)
		}
		_, err = client.Update(&configMap)
		if err != nil {
			return fmt.Errorf("Failed to update ConfigMap %q: %v", configMapName, err)
		}
		fmt.Printf("ConfigMap %q updated\n", configMapName)
	} else {
		fmt.Printf("ConfigMap %q created\n", configMapName)
	}
	return nil
}
