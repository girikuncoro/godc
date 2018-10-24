package libvirt

import (
	"encoding/xml"
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func newDefVolume() libvirtxml.StorageVolume {
	volumeDef := libvirtxml.StorageVolume{
		Target: &libvirtxml.StorageVolumeTarget{
			Format: &libvirtxml.StorageVolumeTargetFormat{
				Type: "qcow2",
			},
			Permissions: &libvirtxml.StorageVolumeTargetPermissions{
				Mode:  "644",
				Owner: "libvirt-qemu",
				Group: "kvm",
			},
		},
		Capacity: &libvirtxml.StorageVolumeSize{
			Unit:  "bytes",
			Value: 1,
		},
	}
	return volumeDef
}

func newDefVolumeFromXML(data string) (libvirtxml.StorageVolume, error) {
	var volDef libvirtxml.StorageVolume
	err := xml.Unmarshal([]byte(data), &volDef)
	if err != nil {
		return libvirtxml.StorageVolume{}, err
	}
	return volDef, nil
}

func newDefVolumeFromLibvirt(vol *libvirt.StorageVol) (libvirtxml.StorageVolume, error) {
	volDefXML, err := vol.GetXMLDesc(0)
	if err != nil {
		return libvirtxml.StorageVolume{}, fmt.Errorf("Not able to get XML description volume: %s", err)
	}

	volDef, err := newDefVolumeFromXML(volDefXML)
	if err != nil {
		return libvirtxml.StorageVolume{}, fmt.Errorf("Not able to get volume definition from XML: %s", err)
	}

	return volDef, nil
}
