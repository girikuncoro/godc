package libvirt

import (
	"encoding/xml"
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

	pool, err := virConn.LookupStoragePoolByName(defaultPoolName)
	if err != nil {
		return fmt.Errorf("Not able to find storage pool %s", pool)
	}
	defer pool.Free()

	volDef := newDefVolume()
	volDef.Name = name

	// TODO(giri): validate image type is qcow2
	// Assume source image is given
	img, err := newImage(source)
	if err != nil {
		return err
	}

	size, err := img.Size()
	if err != nil {
		return err
	}
	log.Printf("Image %s is %d bytes", img, size)
	volDef.Capacity.Unit = "B"
	volDef.Capacity.Value = size

	volDefXML, err := xml.Marshal(volDef)
	if err != nil {
		return fmt.Errorf("Error serializing libvirt volume: %s", err)
	}

	vol, err := pool.StorageVolCreateXML(string(volDefXML), 0)
	if err != nil {
		return fmt.Errorf("Error creating libvirt volume: %s", err)
	}
	defer vol.Free()

	volID, err := vol.GetKey()
	if err != nil {
		return fmt.Errorf("Error retrieving libvirt volume id: %s", err)
	}
	log.Printf("Volume has been created with ID: %s", volID)

	// upload source into the newly created volume
	err = img.Import(newCopier(virConn, vol, volDef.Capacity.Value), volDef)
	if err != nil {
		return fmt.Errorf("Error uploading source %s into new volume: %s", img.String(), err)
	}

	return nil
}
