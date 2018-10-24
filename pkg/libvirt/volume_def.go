package libvirt

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func newDefVolume() libvirtxml.StorageVolume {
	volumeDef := libvirtxml.StorageVolume{
		Target: &libvirtxml.StorageVolumeTarget{
			Format: &libvirtxml.StorageVolumeTargetFormat{
				Type: "qcow2",
			},
			Permissions: &libvirtxml.StorageVolumeTargetPermissions{
				Mode: "644",
			},
		},
		Capacity: &libvirtxml.StorageVolumeSize{
			Unit:  "bytes",
			Value: 1,
		},
	}
	return volumeDef
}
