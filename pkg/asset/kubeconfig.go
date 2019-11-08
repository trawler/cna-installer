package asset

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client function accepts the location of Kubeconfig and returns a Kubernetes Clientset struct
func Client(dir string) (*kubernetes.Clientset, error) {
	kubeconfig, err := getKubeconfig(dir)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return clientset, nil
}

func getKubeconfig(logDir string) (*rest.Config, error) {
	fmt.Printf("full path is: %s/generated/auth/kubeconfig", logDir)
	kubeConfigPath := stringPtr(fmt.Sprintf("%s/generated/auth/kubeconfig", logDir))
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return config, nil
}

// Accepts a string var and return a pointer to the var
func stringPtr(a string) *string {
	return &a
}
