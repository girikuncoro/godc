package libvirt

import (
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

// Config represents libvirt config
type Config struct {
	URI string
}

// Client represents libvirt client
type Client struct {
	libvirt *libvirt.Connect
}

// Client creates new libvirt client
func (c *Config) Client() (*Client, error) {
	libvirtClient, err := libvirt.NewConnect(c.URI)
	if err != nil {
		return nil, err
	}
	log.Println("Created libvirt client")

	client := &Client{
		libvirt: libvirtClient,
	}

	return client, nil
}
