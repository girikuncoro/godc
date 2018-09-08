package vm

import (
	"fmt"

	libvirtc "source.golabs.io/cloud-foundation/godc/pkg/cli/libvirtclient"
)

// VmAPI interface
type VmAPI interface {
	ListVms() error
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
		fmt.Println("ID\tName\t\tUUID")
		fmt.Printf("--------------------------------------------------------\n")
		domains, _ := lc.Domains()
		for _, d := range domains {
			fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
		}
		fmt.Printf("--------------------------------------------------------\n")
	}
	return nil
}
