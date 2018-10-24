package libvirt

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func newVolumeDef() libvirtxml.StorageVolume {
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
			Unit: "bytes",
			// Default to 2GB
			Value: 2147483648,
		},
	}
	return volumeDef
}
