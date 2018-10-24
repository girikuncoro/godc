package libvirt

type Libverter interface {
	DomainCreate() error
}
