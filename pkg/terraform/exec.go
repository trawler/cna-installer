package terraform

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kardianos/osext"
)

// ErrBinaryNotFound is triggered if the TerraForm binary could not be found on disk
var ErrBinaryNotFound = errors.New(
	"TerraForm not in executable's folder, cwd nor PATH",
)

const (
	logsFolderName = "logs"
)

// Executor enables calling TerraForm from Go, across platforms, with any
// additional providers/provisioners that the currently executing binary
// exposes.
//
// The TerraForm binary is expected to be in the executing binary's folder, in
// the current working directory or in the PATH.
// Each Executor runs in a temporary folder, so each Executor should only be
// used for one TF project.
//
// Between the unreliability of the internal interfaces in the terraform library and then
// need to communicate with providers, we'll wrap the terraform command in bash, rather
// than importing the `github.com/hashicorp/terraform` library and calling methods
// directly. See https://github.com/hashicorp/terraform/issues/12582 for more info.
type Executor struct {
	binaryPath       string
	version          string
	workingDirectory string
	envVariables     map[string]string
}

// NewTerraformClient return a struct which behaves like the cli terraform client.
func NewTerraformClient(workingDirectory string) (*Executor, error) {
	ex := new(Executor)
	ex.workingDirectory = workingDirectory

	// Find the TerraForm binary.
	out, err := tfBinaryPath()
	if err != nil {
		return nil, err
	}
	ex.binaryPath = out
	return ex, nil
}

// Init runs "terraform init" action
func (cli *Executor) Init(params *TfInitParams) *TfAction {
	return &TfAction{
		action: "init",
		bin:    cli,
		Dir:    cli.workingDirectory,
		params: params,
	}
}

// Plan runs "terraform plan" action
func (cli *Executor) Plan(params *TfPlanParams) *TfAction {
	return &TfAction{
		action: "plan",
		bin:    cli,
		Dir:    cli.workingDirectory,
		params: params,
	}
}

// Apply runs "terraform apply" action
func (cli *Executor) Apply(params *TfPlanParams) *TfAction {
	// set auto-approve to true for apply actions
	if params.AutoApprove == false {
		params.AutoApprove = true
	}
	return &TfAction{
		action: "apply",
		bin:    cli,
		Dir:    cli.workingDirectory,
		params: params,
	}
}

func (cli *Executor) fetchVersion() {
	cli.version = "dev"
}

// tfBinatyPath searches for a TerraForm binary on disk:
// - in the executing binary's folder,
// - in the current working directory,
// - in the PATH.
// The first to be found is the one returned.
func tfBinaryPath() (string, error) {
	// Depending on the platform, the expected binary name is different.
	binaryFileName := "terraform"

	// Look into the executable's folder.
	if execFolderPath, err := osext.ExecutableFolder(); err == nil {
		path := filepath.Join(execFolderPath, binaryFileName)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			return path, nil
		}
	}

	// Look into cwd.
	if workingDirectory, err := os.Getwd(); err == nil {
		path := filepath.Join(workingDirectory, binaryFileName)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			return path, nil
		}
	}

	// If we still haven't found the executable, look for it
	// in the PATH.
	if path, err := exec.LookPath(binaryFileName); err == nil {
		return filepath.Abs(path)
	}

	return "", ErrBinaryNotFound
}
