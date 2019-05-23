package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/cli/vm"
)

// getVmPre is prerunner of getVmCmd
func getVmPre(c *Cli) error {
	// validate host endpoints
	if len(c.v.GetStringSlice(configKeyHosts)) == 0 {
		return fmt.Errorf("Host endpoints are not provided")
	}
	return nil
}

// getVmRun is runner of getVmCmd
func getVmRun(c *Cli) error {
	vc, err := vm.NewVMClient(c.v.GetStringSlice(configKeyHosts))
	if err != nil {
		fmt.Printf("Getting vm error: %s", err.Error())
		return err
	}

	vc.ListVms()
	return nil
}
