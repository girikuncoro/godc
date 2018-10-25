package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type Metricer interface {
	Usage() uint64
}

type Metrics struct {
	Memory Metricer
}

type Memory struct {
	Used  uint64
	Free  uint64
	Total uint64
}

func (m *Memory) Usage() uint64 {
	p := 100 * m.Used / m.Total
	return p
}

func Collect() (*Metrics, error) {
	memory, err := collectMemory()
	if err != nil {
		return nil, err
	}
	return &Metrics{
		Memory: memory,
	}, nil
}

func collectMemory() (*Memory, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("Error retrieving virtual memory")
	}
	return &Memory{
		Used:  vmem.Used,
		Free:  vmem.Free,
		Total: vmem.Total,
	}, nil
}
