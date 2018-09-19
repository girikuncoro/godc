package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/cli/vm"
	"github.com/girikuncoro/godc/pkg/config"
)

// listVolumePre is prerunner of listVolumeCmd
func listVolumePre(c *Cli) error {
	// validate host endpoints
	if len(c.hostEndpoints) == 0 {
		return fmt.Errorf("Host endpoints are not provided")
	}
	return nil
}

// listVolumeRun is runner of listVolumeCmd
func listVolumeRun(c *Cli) error {
	vc, err := vm.NewVMClient(c.hostEndpoints)
	if err != nil {
		fmt.Printf("Listing vm error: %s", err.Error())
		return err
	}

	vc.ListVolumes(config.DefaultStoragePool)
	return nil
}
