package terraform

import "strconv"

// TfPlanParams comment
type TfPlanParams struct {
	Destroy          bool
	DetailedExitcode bool

	Input       *bool
	Lock        *bool
	LockTimeout int
	NoColor     bool

	ModuleDepth *int
	Out         *string
	Parallelism *int
	Refresh     *bool
	State       *string
	Target      []*string
	Var         map[string]string
	VarFile     []*string
}

// NewTfPlanParams comment
func NewTfPlanParams() *TfPlanParams {
	return &TfPlanParams{}
}

// Opts comment
func (p *TfPlanParams) Opts() map[string][]string {
	opts := make(map[string][]string)

	if p.Destroy {
		opts["destroy"] = []string{""}
	}

	if p.DetailedExitcode {
		opts["detailed-exitcode"] = []string{""}
	}

	if p.Input != nil && *p.Input == false {
		opts["input"] = []string{"false"}
	}

	if p.Lock != nil {
		if *p.Lock == true {
			opts["lock"] = []string{"true"}
		} else {
			opts["lock"] = []string{"false"}
		}
	}

	if p.LockTimeout != 0 {
		opts["lock-timeout"] = []string{strconv.Itoa(p.LockTimeout) + "s"}
	}

	if p.ModuleDepth != nil {
		opts["module-depth"] = []string{strconv.Itoa(*p.ModuleDepth)}
	}

	if p.NoColor {
		opts["no-color"] = []string{""}
	}

	if p.Out != nil {
		opts["out"] = []string{*p.Out}
	}

	if p.Parallelism != nil {
		opts["parallelism"] = []string{strconv.Itoa(*p.Parallelism)}
	}

	if p.Refresh != nil {
		opts["refresh"] = []string{"false"}
	}

	if p.State != nil {
		opts["state"] = []string{*p.State}
	}

	if p.Target != nil {
		opts["target"] = StringSlice(p.Target)
	}

	if p.Var != nil {
		vars := []string{}
		for key, value := range p.Var {
			vars = append(vars, key+"="+value)
		}
		opts["var"] = vars
	}

	if p.VarFile != nil {
		opts["var-file"] = StringSlice(p.VarFile)
	}

	return opts
}

// OptsString comment
func (p *TfPlanParams) OptsString() string {
	return extractOptsString(p)
}

// OptsStringSlice comment
func (p *TfPlanParams) OptsStringSlice() []string {
	return extractOptsStringSlice(p)
}
