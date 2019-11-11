package assets

import (
	"fmt"

	"github.com/trawler/cna-installer/pkg/assets/manifests"
	"k8s.io/client-go/kubernetes"
)

// NewClient function accepts the location of Kubeconfig and returns a Kubernetes Clientset struct
func NewClient(kubeconfigPath string) (*kubernetes.Clientset, error) {
	kubeconfig, err := getKubeconfig(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return clientset, nil
}

// Install deploys all the required assets onto the cluster
func Install(client *kubernetes.Clientset) {
	manifests.CreateNamespace(client, "cna-installer")
	manifests.InstallTraefikIngressController(client)
}