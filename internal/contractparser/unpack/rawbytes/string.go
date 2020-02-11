package rawbytes

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type stringDecoder struct{}

// Decode -
func (d stringDecoder) Decode(dec io.Reader, code *strings.Builder) (int, error) {
	b := make([]byte, 4)
	if n, err := dec.Read(b); err != nil {
		return n, err
	}

	length := int(binary.BigEndian.Uint32(b))
	data := make([]byte, length)
	if _, err := dec.Read(data); err != nil && err != io.EOF {
		return 4 + length, err
	}
	fmt.Fprintf(code, `{ "string": "%s" }`, data)
	return 4 + length, nil
}
