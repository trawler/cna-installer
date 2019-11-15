package utils

import (
	"fmt"

	"k8s.io/client-go/kubernetes"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateDeployment comment
func CreateDeployment(k8sClient *kubernetes.Clientset,
	appName string,
	deploymentName string,
	serviceAccountName string,
	cspec []ContainerSpec,
	namespace string) error {

	labels := map[string]string{
		"app": appName,
	}

	containerSpec, err := generateContainerSpec(cspec)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deploymentName,
					Namespace: namespace,
					Labels:    labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName:            serviceAccountName,
					TerminationGracePeriodSeconds: int64Ptr(60),
					Containers:                    containerSpec,
				},
			},
		},
	}

	client := k8sClient.AppsV1().Deployments(namespace)
	_, err = client.Create(&deployment)
	if err != nil {
		if !apierr.IsAlreadyExists(err) {
			return fmt.Errorf("Failed to create Deployment %s: %v", deploymentName, err)
		}
		_, err = client.Update(&deployment)
		if err != nil {
			return fmt.Errorf("Failed to update Deployment %q: %v", deploymentName, err)
		}
		fmt.Printf("Deployment %q updated\n", deploymentName)
	} else {
		fmt.Printf("Deployment %q created\n", deploymentName)
	}
	return nil
}

// helper functions
func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }
