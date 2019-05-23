package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/cli/vm"
	"github.com/girikuncoro/godc/pkg/config"
)

// getVolumePre is prerunner of getVolumeCmd
func getVolumePre(c *Cli) error {
	// validate host endpoints
	if len(c.v.GetStringSlice(configKeyHosts)) == 0 {
		return fmt.Errorf("Host endpoints are not provided")
	}
	return nil
}

// getVolumeRun is runner of getVolumeCmd
func getVolumeRun(c *Cli) error {
	vc, err := vm.NewVMClient(c.v.GetStringSlice(configKeyHosts))
	if err != nil {
		fmt.Printf("Getting volume error: %s", err.Error())
		return err
	}

	vc.ListVolumes(config.DefaultStoragePool)
	return nil
}
