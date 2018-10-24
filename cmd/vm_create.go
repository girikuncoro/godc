package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/libvirt"
)

// createVmPre is prerunner of createVmCmd
func createVmPre(c *Cli) error {
	if len(c.hostEndpoints) != 1 {
		return fmt.Errorf("single hostendpoint must be provided")
	}

	if c.vmCmd.name == "" {
		return fmt.Errorf("volume name must be provided")
	}

	return nil
}

// createVmRun is runner of createVmCmd
func createVmRun(c *Cli) error {
	config := libvirt.Config{
		URI: c.hostEndpoints[0],
	}

	client, err := config.Client()
	if err != nil {
		return err
	}

	err = libvirt.DomainCreate(client, c.vmCmd.name)
	if err != nil {
		return err
	}

	return nil
}
