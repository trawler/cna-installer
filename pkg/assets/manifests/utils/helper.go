package utils

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// generateRbacRules takes a generic structure and returns it
// as a a PolicyRule array
func generateRbacRules(rules []PolicyRule) ([]rbacv1.PolicyRule, error) {
	var result []rbacv1.PolicyRule
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if err = decoder.Decode(rules); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return result, nil
}

// generateContainerSpec takes a generic structure and returns it
// as a corev1.Container array
func generateContainerSpec(containerSpec []ContainerSpec) ([]corev1.Container, error) {
	var result []corev1.Container

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if err = decoder.Decode(containerSpec); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return result, nil
}

// ParseConfigData accepts a yaml string input and outputs it in a map[string][string] format
func ParseConfigData(config string) (*ConfigMapData, error) {
	t := ConfigMapData{}

	err := yaml.Unmarshal([]byte(config), &t)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &t, nil
}
