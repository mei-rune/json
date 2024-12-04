//go:build !go1.23
// +build !go1.23

package json

import (
	"io"

	"encoding/json"
)

func UnmarshalV2WithNumber(in io.Reader, value interface{}) error {
	decoder := json.NewDecoder(in)
	decoder.UseNumber()
	return decoder.Decode(value)
}

func UnmarshalBytesV2WithNumber(bs []byte, value interface{}) error {
	return UnmarshalWithJsonNumber(bs, &value)
}
