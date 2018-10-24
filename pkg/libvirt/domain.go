package libvirt

import (
	"fmt"
	"log"
)

// DomainCreate create new domain via libvirt
func DomainCreate(client *Client, name string) error {
	log.Printf("Libvirt domain create")

	virConn := client.libvirt
	if virConn == nil {
		return fmt.Errorf("Libvirt virConn is nil")
	}

	domainDef := newDomainDef()
	domainDef.Name = name
	setConsoles(&domainDef)

	connectURI, err := virConn.GetURI()
	if err != nil {
		return fmt.Errorf("Error retrieving libvirt connection URI: %s", err)
	}
	log.Printf("Creating libvirt domain at %s", connectURI)

	data, err := xmlMarshallIndented(domainDef)
	if err != nil {
		return fmt.Errorf("Error serializing libvirt domain: %s", err)
	}

	log.Printf("Creating libvirt domain with XML:\n%s", data)

	domain, err := virConn.DomainDefineXML(data)
	if err != nil {
		return fmt.Errorf("Error defining libvirt domain: %s", err)
	}

	err = domain.Create()
	if err != nil {
		return fmt.Errorf("Error creating libvirt domain: %s", err)
	}
	defer domain.Free()

	id, err := domain.GetUUIDString()
	if err != nil {
		return fmt.Errorf("Error retrieving libvirt domain id: %s", err)
	}

	log.Printf("Libvirt domain has been created with id: %s", id)
	return nil
}
