package libvirt

import (
	"fmt"

	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func newDefDisk(i int) libvirtxml.DomainDisk {
	diskDef := libvirtxml.DomainDisk{
		Device: "disk",
		Target: &libvirtxml.DomainDiskTarget{
			Bus: "virtio",
			Dev: fmt.Sprintf("vd%s", diskLetterForIndex(i)),
		},
		Driver: &libvirtxml.DomainDiskDriver{
			Name: "qemu",
			Type: "raw",
		},
	}
	return diskDef
}
