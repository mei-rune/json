//go:build go1.23
// +build go1.23

package json

import (
	"io"

	jsonv2 "github.com/go-json-experiment/json"
	// json "github.com/goccy/go-json"
)

func UnmarshalV2WithNumber(in io.Reader, value interface{}) error {
	return jsonv2.UnmarshalRead(in, &value, jsonv2.StringifyNumbers(true))
}

func UnmarshalBytesV2WithNumber(bs []byte, value interface{}) error {
	return jsonv2.Unmarshal(bs, &value, jsonv2.StringifyNumbers(true))
}
