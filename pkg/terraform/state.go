package terraform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/terraform/states/statefile"
	"github.com/pkg/errors"
)

// State Struct
type State struct {
	Resources []StateResource `json:"resources"`
}

// StateResource is local sparse representation of terraform state resource
type StateResource struct {
	Module    string                  `json:"module"`
	Name      string                  `json:"name"`
	Type      string                  `json:"type"`
	Instances []StateResourceInstance `json:"instances"`
}

// StateResourceInstance is an instance of terraform state resource.
type StateResourceInstance struct {
	Attributes map[string]interface{} `json:"attributes"`
}

// ErrResourceNotFound is an error that instructs that requested resource was not found.
var ErrResourceNotFound = fmt.Errorf("resource not found")

// LookupResource finds a resource for a given module, type and name from the state.
// If module is "root", it is treated as ""
// If no resource is found for the triplet, ErrResourceNotFound error is returned.
func LookupResource(state *State, module, t, name string) (*StateResource, error) {
	if module == "root" {
		module = ""
	}
	for idx, r := range state.Resources {
		if module == r.Module && t == r.Type && name == r.Name {
			return &state.Resources[idx], nil
		}
	}
	return nil, ErrResourceNotFound
}

// ReadStateFile returns that terraform state from the file.
func ReadStateFile(file string) (*State, error) {
	sfRaw, err := ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %q", file)
	}

	var tfstate State
	if err := json.Unmarshal(sfRaw, &tfstate); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %q", file)
	}
	if len(tfstate.Resources) == 0 {
		return nil, fmt.Errorf("found empty state file: %v", file)
	}

	return &tfstate, nil
}

// ReadFile reads the terraform state from file and returns the contents in bytes
// It returns an error if reading the state was unsuccessful
// ReadState utilizes the terraform's internal wiring to upconvert versions of terraform state to return
// the state it currently recognizes.
func ReadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %q", file)
	}
	defer f.Close()

	sf, err := statefile.Read(f)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read statefile from %q", file)
	}

	out := bytes.Buffer{}
	if err := statefile.Write(sf, &out); err != nil {
		return nil, errors.Wrapf(err, "failed to write statefile")
	}
	return out.Bytes(), nil
}
