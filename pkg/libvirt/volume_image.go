package libvirt

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

type image interface {
	Size() (uint64, error)
	String() string
	Import(func(io.Reader) error, libvirtxml.StorageVolume) error
}

func newImage(source string) (image, error) {
	url, err := url.Parse(source)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse source '%s' as url: %s", source, err)
	}
	if strings.HasPrefix(url.Scheme, "http") {
		return &httpImage{url: url}, nil
	} else {
		return nil, fmt.Errorf("Not able to read from '%s': %s", url.String(), err)
	}
}

type httpImage struct {
	url *url.URL
}

func (i *httpImage) String() string {
	return i.url.String()
}

func (i *httpImage) Size() (uint64, error) {
	resp, err := http.Head(i.url.String())
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0,
			fmt.Errorf(
				"Error accessing remote source: %s - %s",
				i.url.String(),
				resp.Status,
			)
	}

	length, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		err = fmt.Errorf(
			"Error getting Content-Length of %s : %s - got %s",
			i.url.String(),
			err,
			resp.Header.Get("Content-Length"),
		)
		return 0, err
	}
	return uint64(length), nil
}

func (i *httpImage) Import(copier func(io.Reader) error, vol libvirtxml.StorageVolume) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", i.url.String(), nil)
	if err != nil {
		return fmt.Errorf("Error downloading %s: %s", i.url.String(), err)
	}

	// TODO(giri): must check if we have downloaded this image before

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error downloading %s: %s", i.url.String(), err)
	}

	defer resp.Body.Close()
	log.Printf("Url response status code %s\n", resp.Status)
	return copier(resp.Body)
}
