package libvirtclient

import (
	"fmt"
	"net"
	"time"

	libvirt "github.com/digitalocean/go-libvirt"
)

// Libverter implements libvert caller
type Libverter interface {
	HostEndpoint() string
	Domains() ([]libvirt.Domain, error)
	Volumes(string, int) ([]string, error)
}

// LibvirtClient holds libvirt caller and host endpoint
type LibvirtClient struct {
	libvirtc     *libvirt.Libvirt
	hostEndpoint string
}

// GetClient get Libvirt client
func GetClient(hostEndpoint string) (Libverter, error) {
	if hostEndpoint == "" {
		return nil, fmt.Errorf("hostEndpoint is empty")
	}

	c, err := net.DialTimeout("tcp", hostEndpoint, 2*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to dial libvirt: %v", err)
	}

	l := &LibvirtClient{
		libvirtc:     libvirt.New(c),
		hostEndpoint: hostEndpoint,
	}
	if err := l.libvirtc.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	return l, nil
}

// HostEndpoint returns hostEndpoint
func (lc *LibvirtClient) HostEndpoint() string {
	return lc.hostEndpoint
}

// Domains returns list of VM domains
func (lc *LibvirtClient) Domains() ([]libvirt.Domain, error) {
	domains, err := lc.libvirtc.Domains()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve domains: %v", err)
	}

	return domains, nil
}

// Volumes returns list of volumes from specified storage pool
func (lc *LibvirtClient) Volumes(poolName string, maxList int) ([]string, error) {
	storagePool, err := lc.libvirtc.StoragePool(poolName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve storage pool: %v", err)
	}

	volumes, err := lc.libvirtc.StoragePoolListVolumes(storagePool, int32(maxList))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve volumes: %v", err)
	}

	return volumes, nil
}
