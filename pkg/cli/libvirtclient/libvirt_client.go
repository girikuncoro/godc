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

	for _, lll := range domains {
		ll, _ := lc.libvirtc.DomainInterfaceAddresses(lll, 0, 0)
		fmt.Println(ll)

		l1, l2, l3, l4, l5, _ := lc.libvirtc.DomainGetInfo(lll)
		fmt.Printf("stat: %d", l1)
		fmt.Printf(" maxmem: %d", l2)
		fmt.Printf(" mem: %d", l3)
		fmt.Printf(" cpu: %d", l4)
		fmt.Printf(" cputime: %d\n", l5)

		lc.libvirtc.DomainGetXMLDesc(lll, lc.libvirtc.DomainDefineXMLFlags("interface"))
	}

	return domains, nil
}
