package main

import (
	"fmt"
	"os/exec"

	"github.com/kubernetes/pkg/kubelet/kubeletconfig/util/log"
)

type Monitor interface {
	Active()
}

type ServiceMonitor struct {
	serviceName string
}

func NewServiceMonitor(serviceName string) Monitor {
	return &ServiceMonitor{serviceName}
}

func (sm *ServiceMonitor) Active() {
	cmd := exec.Command("systemctl", "status", sm.serviceName)
	out, err := cmd.Output()
	if err != nil {
		log.Errorf("Error executing systemctl: %v", err)
	}
	fmt.Printf("%s\n\n", out)
}
