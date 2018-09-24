package main

import (
	"log"
	"os/exec"
	"strings"
)

const runStatus = "active (running)"

type Monitor interface {
	Active() bool
}

type ServiceMonitor struct {
	serviceName string
}

func NewServiceMonitor(serviceName string) Monitor {
	return &ServiceMonitor{serviceName}
}

func (sm *ServiceMonitor) Active() bool {
	cmd := exec.Command("systemctl", "status", sm.serviceName)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing systemctl: %v", err)
		return false
	}

	return strings.Contains(string(out), runStatus)
}
