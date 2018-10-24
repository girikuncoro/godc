package libvirt

import (
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

type Config struct {
	URI string
}

type Client struct {
	libvirt *libvirt.Connect
}

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
