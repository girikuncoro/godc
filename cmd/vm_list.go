package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/cli/vm"
)

// listVmPre is prerunner of listVmCmd
func listVmPre(c *Cli) error {
	// validate host endpoints
	if len(c.v.GetStringSlice(configKeyHosts)) == 0 {
		return fmt.Errorf("Host endpoints are not provided")
	}
	return nil
}

// listVmRun is runner of listVmCmd
func listVmRun(c *Cli) error {
	vc, err := vm.NewVMClient(c.v.GetStringSlice(configKeyHosts))
	if err != nil {
		fmt.Printf("Listing vm error: %s", err.Error())
		return err
	}

	vc.ListVms()
	return nil
}
