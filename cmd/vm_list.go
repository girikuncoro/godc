package main

import (
	"fmt"

	"source.golabs.io/cloud-foundation/godc/pkg/cli/vm"
)

// listVmPre is prerunner of listVmCmd
func listVmPre(c *Cli) error {
	// validate host endpoints
	if len(c.hostEndpoints) == 0 {
		return fmt.Errorf("Host endpoints are not provided")
	}
	return nil
}

// listVmRun is runner of listVmCmd
func listVmRun(c *Cli) error {
	vc, err := vm.NewVMClient(c.hostEndpoints)
	if err != nil {
		fmt.Printf("Listing vm error: %s", err.Error())
		return err
	}

	vc.ListVms()
	return nil
}
