package libvirt

import (
	"fmt"
	"log"
)

const (
	defaultPoolName = "default"
)

// VolumeCreate creates new volume
func VolumeCreate(client *Client, name string, source string) error {
	log.Printf("Libvirt volume create")

	virConn := client.libvirt
	if virConn == nil {
		return fmt.Errorf("Libvirt virConn is nil")
	}

	volumeDef := newVolumeDef()
	volumeDef.Name = defaultPoolName

	// TODO(giri): validate image type is qcow2
	// Assume source image is given
	_, err := newImage(source)
	if err != nil {
		return err
	}
	return nil
}
