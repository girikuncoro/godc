package libvirt

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func xmlMarshallIndented(b interface{}) (string, error) {
	buf := new(bytes.Buffer)
	enc := xml.NewEncoder(buf)
	enc.Indent("  ", "    ")
	if err := enc.Encode(b); err != nil {
		return "", fmt.Errorf("could not marshall +%v", b)
	}
	return buf.String(), nil
}
