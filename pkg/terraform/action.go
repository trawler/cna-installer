package terraform

import (
	"os"
	"os/exec"
	"sort"
	"strings"
)

// TfActionParams comment
type TfActionParams interface {
	Opts() map[string][]string
	OptsString() string
	OptsStringSlice() []string
}

// TfAction comment
type TfAction struct {
	Cmd    *exec.Cmd
	Dir    string
	action string
	bin    *Executor
	opts   map[string]string
	params TfActionParams
}

// Initialise comment
func (a *TfAction) Initialise() *TfAction {
	args := append([]string{a.action}, a.params.OptsStringSlice()...)
	a.Cmd = exec.Command(a.bin.binaryPath, args...)
	if a.Dir != "" {
		a.Cmd.Dir = a.Dir
	}

	a.Cmd.Stdout = os.Stdout
	a.Cmd.Stderr = os.Stderr

	return a
}

// Run the terraform command
func (a *TfAction) Run() (err error) {
	a.Cmd.Start()
	return a.Cmd.Wait()
}

// BoolPtr comment
func BoolPtr(a bool) *bool {
	return &a
}

// TruePtr comment
func TruePtr() *bool {
	return BoolPtr(true)
}

// FalsePtr comment
func FalsePtr() *bool {
	return BoolPtr(false)
}

// StringPtr comment
func StringPtr(a string) *string {
	return &a
}

// IntPtr comment
func IntPtr(a int) *int {
	return &a
}

// StringSlice comment
func StringSlice(a []*string) (o []string) {
	o = make([]string, len(a))
	for i, e := range a {
		o[i] = *e
	}
	return
}

// StringMapPtr comment
func StringMapPtr(a map[string]string) *map[string]string {
	return &a
}

// extractOptsStringSlice comment
func extractOptsStringSlice(p TfActionParams) (options []string) {
	opts := p.Opts()
	keys := mapStringSliceKeys(opts)
	sort.Strings(keys)

	outputs := make([]string, 0)
	for _, key := range keys {
		value := opts[key]
		sort.Strings(value)
		for _, val := range value {
			output := "-" + key
			if val != "" {
				switch key {
				case "var":
					outputs = append(outputs, output)
					outputs = append(outputs, "'"+val+"'")
					continue
				default:
					output = output + "=" + val
				}
			}
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func extractOptsString(p TfActionParams) (options string) {
	return strings.Join(
		extractOptsStringSlice(p),
		" ",
	)
}

func mapStringSliceKeys(s map[string][]string) (keys []string) {
	keys = make([]string, len(s))

	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return
}
