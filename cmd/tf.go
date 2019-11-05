package cmd

import (
	"fmt"

	"github.com/trawler/cna-installer/pkg/terraform"
)

func tfRun() error {
	// Populate TF_VAR Environment variables
	if err = terraform.GetEnvVars(cluster); err != nil {
		return fmt.Errorf("%v", err)
	}

	// Get Init opts
	initParams.Opts()

	// Run terraform init
	init := tf.Init(initParams)

	init.Initialise()
	init.Run()

	// Run terraform plan
	planParams.Opts()
	plan := tf.Plan(planParams)
	plan.Initialise()

	if err = plan.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	// // Run terraform apply
	// apply := tf.Apply(planParams)
	// apply.Initialise()
	//
	// if err = apply.Run(); err != nil {
	// 	return fmt.Errorf("%v", err)
	// }
	//
	return nil
}

func tfDestroy() error {
	// Populate TF_VAR Environment variables
	if err = terraform.GetEnvVars(cluster); err != nil {
		return fmt.Errorf("%v", err)
	}

	destroy := tf.Destroy(planParams)
	destroy.Initialise()

	if err = destroy.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
