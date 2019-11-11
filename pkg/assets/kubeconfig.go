package assets

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeconfig(logDir string) (*rest.Config, error) {
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
