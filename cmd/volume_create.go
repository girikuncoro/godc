package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/libvirt"
)

// createVolumePre is prerunner of createVolumeCmd
func createVolumePre(c *Cli) error {
	// validate host endpoints
	if len(c.hostEndpoints) != 1 {
		return fmt.Errorf("Single hostEndpoint must be provided")
	}

	return nil
}

// createVolumeRun is runner of createVolumeCmd
func createVolumeRun(c *Cli) error {
	config := libvirt.Config{
		URI: c.hostEndpoints[0],
	}

	client, err := config.Client()
	if err != nil {
		return err
	}

	// TODO(giri): To be implemented with libvirt package
	fmt.Println("Create volume is not implemented yet")
	fmt.Printf("Client: +%v", client)

	return nil
}
