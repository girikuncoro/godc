package main

import (
	"fmt"

	"github.com/girikuncoro/godc/pkg/libvirt"
)

// createVolumePre is prerunner of createVolumeCmd
func createVolumePre(c *Cli) error {
	if len(c.v.GetStringSlice(configKeyHosts)) != 1 {
		return fmt.Errorf("single hostendpoint must be provided")
	}

	if c.volumeCmd.name == "" {
		return fmt.Errorf("volume name must be provided")
	}

	if c.volumeCmd.source == "" {
		return fmt.Errorf("volume source url must be provided")
	}

	return nil
}

// createVolumeRun is runner of createVolumeCmd
func createVolumeRun(c *Cli) error {
	config := libvirt.Config{
		URI: c.v.GetStringSlice(configKeyHosts)[0],
	}

	client, err := config.Client()
	if err != nil {
		return err
	}

	err = libvirt.VolumeCreate(client, c.volumeCmd.name, c.volumeCmd.source)
	if err != nil {
		return err
	}

	return nil
}
