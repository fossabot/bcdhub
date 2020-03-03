package rawbytes

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/aopoltorzhicky/bcdhub/internal/contractparser/unpack/tzbase58"
)

type bytesDecoder struct{}

// Decode -
func (d bytesDecoder) Decode(dec *decoder, code *strings.Builder) (int, error) {
	length, err := decodeLength(dec)
	if err != nil {
		return 4, err
	}
	if dec.Len() < length {
		return 4, &invalidDataError{
			typ:     "bytes",
			message: fmt.Sprintf("Not enough data in reader: %d < %d", dec.Len(), length),
		}
	}

	data := make([]byte, length)
	if _, err := dec.Read(data); err != nil && err != io.EOF {
		return 4 + length, err
	}

	// log.Printf("[bytes Decode] data: %x\n", data)

	if length == tzbase58.KeyHashLength {
		if res, err := decodeKeyHash(data); err == nil {
			fmt.Fprintf(code, `{ "string": "%s" }`, res)
			return 4 + length, nil
		}
	}

	if length == tzbase58.AddressLength {
		if res, err := decodeAddress(data); err == nil {
			fmt.Fprintf(code, `{ "string": "%s" }`, res)
			return 4 + length, nil
		}
	}

	s := hex.EncodeToString(data)
	intDec := newDecoder(strings.NewReader(s))

	var bufBuilder strings.Builder
	l, err := hexToMicheline(intDec, &bufBuilder)
	if err != nil || intDec.Len() > 0 {
		if _, ok := err.(*invalidDataError); ok || intDec.Len() > 0 {
			fmt.Fprintf(code, `{ "bytes": "%x" }`, data)
			return 4 + length, nil
		}
		return l + 4, err
	}
	fmt.Fprintf(code, `{ "bytes": "%s" }`, bufBuilder.String())
	return 4 + length, nil
}

func decodeKeyHash(data []byte) (string, error) {
	return tzbase58.DecodeKeyHash(hex.EncodeToString(data))
}

func decodeAddress(data []byte) (string, error) {
	if tzbase58.HasKT1Affixes(data) {
		if res, err := tzbase58.DecodeKT(hex.EncodeToString(data)); err == nil {
			return res, nil
		}
	}

	if res, err := tzbase58.DecodeTz(hex.EncodeToString(data)); err == nil {
		return res, nil
	}

	return "", fmt.Errorf("decodeAddress: can't decode address from bytes: %v", data)
}
