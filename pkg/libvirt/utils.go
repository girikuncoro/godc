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

var diskLetters = []rune("abcdefghijklmnopqrstuvwxyz")

func diskLetterForIndex(idx int) string {
	m := idx / len(diskLetters)
	n := idx % len(diskLetters)
	letter := diskLetters[n]

	if m == 0 {
		return fmt.Sprintf("%c", letter)
	}
	return fmt.Sprintf("%s%c", diskLetterForIndex(m-1), letter)
}
