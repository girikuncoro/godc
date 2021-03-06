package vm

import (
	"fmt"

	libvirtc "github.com/girikuncoro/godc/pkg/cli/libvirtclient"
	"github.com/girikuncoro/godc/pkg/config"
)

// VmAPI interface
type VmAPI interface {
	ListVms() error
	ListVolumes(string) error
}

// Client for sending request.
// Currently it's calling each host machine directly
// as our model is client-only without server. Libvirt in
// host machine is expected to have TCP listen enabled.
type vmClient struct {
	libvirtClients []libvirtc.Libverter
}

// NewVMClient return a vm client
func NewVMClient(hostEndpoints []string) (VmAPI, error) {
	if len(hostEndpoints) == 0 {
		return nil, fmt.Errorf("List of hostEndpoints is empty")
	}

	vc := &vmClient{}
	for _, he := range hostEndpoints {
		lc, err := libvirtc.GetClient(he)
		if err != nil {
			return nil, fmt.Errorf("error getting libvirt client: %v", err)
		}
		vc.libvirtClients = append(vc.libvirtClients, lc)
	}
	return vc, nil
}

// ListVms list vms running in all host machines
func (vc *vmClient) ListVms() error {
	for _, lc := range vc.libvirtClients {
		fmt.Printf("\nHost Endpoint: %s\n\n", lc.HostEndpoint())
		fmt.Println("ID\tName")
		fmt.Printf("--------------------------------------------------------\n")
		domains, _ := lc.Domains()
		for _, d := range domains {
			fmt.Printf("%d\t%s\n", d.ID, d.Name)
		}
		fmt.Printf("--------------------------------------------------------\n")
	}
	return nil
}

// ListVolumes list volumes stored in specified pool
func (vc *vmClient) ListVolumes(poolStorage string) error {
	for _, lc := range vc.libvirtClients {
		fmt.Printf("\nHost Endpoint: %s\n\n", lc.HostEndpoint())
		fmt.Println("Volume")
		fmt.Printf("--------------------------------------------------------\n")
		volumes, _ := lc.Volumes(config.DefaultStoragePool, config.MaxVolumeListPerPool)
		for _, v := range volumes {
			fmt.Printf("%s\n", v)
		}
		fmt.Printf("--------------------------------------------------------\n")
	}
	return nil
}
