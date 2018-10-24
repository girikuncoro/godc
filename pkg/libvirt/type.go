package libvirt

// Libverter represents libvirt functionalities
type Libverter interface {
	DomainCreate() error
}
