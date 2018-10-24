package libvirt

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func newDomainDef() libvirtxml.Domain {
	domainDef := libvirtxml.Domain{
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Type: "hvm",
			},
		},
		Memory: &libvirtxml.DomainMemory{
			Unit:  "MiB",
			Value: 512,
		},
		VCPU: &libvirtxml.DomainVCPU{
			Placement: "static",
			Value:     1,
		},
		CPU: &libvirtxml.DomainCPU{},
		Devices: &libvirtxml.DomainDeviceList{
			Graphics: []libvirtxml.DomainGraphic{
				{
					Spice: &libvirtxml.DomainGraphicSpice{
						AutoPort: "yes",
					},
				},
			},
			Channels: []libvirtxml.DomainChannel{
				{
					Target: &libvirtxml.DomainChannelTarget{
						VirtIO: &libvirtxml.DomainChannelTargetVirtIO{
							Name: "org.qemu.guest_agent.0",
						},
					},
				},
			},
			RNGs: []libvirtxml.DomainRNG{
				{
					Model: "virtio",
					Backend: &libvirtxml.DomainRNGBackend{
						Random: &libvirtxml.DomainRNGBackendRandom{},
					},
				},
			},
		},
		Features: &libvirtxml.DomainFeatureList{
			PAE:  &libvirtxml.DomainFeature{},
			ACPI: &libvirtxml.DomainFeature{},
			APIC: &libvirtxml.DomainFeatureAPIC{},
		},
	}

	domainDef.Type = "kvm"
	return domainDef
}

// setConsoles hardcode console config of domain to serial and virtio
func setConsoles(domainDef *libvirtxml.Domain) {
	consoles := []libvirtxml.DomainConsole{
		libvirtxml.DomainConsole{
			Target: &libvirtxml.DomainConsoleTarget{
				Type: "serial",
			},
		},
		libvirtxml.DomainConsole{
			Target: &libvirtxml.DomainConsoleTarget{
				Type: "virtio",
			},
		},
	}

	for i, console := range consoles {
		port := uint(i)
		console.Target.Port = &port
		domainDef.Devices.Consoles = append(domainDef.Devices.Consoles, console)
	}
}
